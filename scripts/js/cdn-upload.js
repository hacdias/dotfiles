#!/usr/bin/env node
'use strict'

import { put } from 'got'
import meow from 'meow'
import { createHash } from 'crypto'
import { extname, basename } from 'path'
import { readFile } from 'fs-extra'
import sharp from 'sharp'

const config = {
  zone: process.env.BUNNY_ZONE,
  key: process.env.BUNNY_KEY
}

const typeOpts = {
  jpeg: {
    progressive: true
  }
}

const presets = {
  images: {
    jpeg: [
      [250],
      [500],
      [1000],
      [2000]
    ],
    webp: [
      [250],
      [500],
      [1000],
      [2000]
    ]
  }
}

const presetsPath = {
  images: 'i',
  default: 'u'
}

const cli = meow(`
  Uploads to BunnyCDN. The preset will not only determine the transformations
  done to the image (if it's an image), but also to the path where they'll be
  uploaded to.

  Preset books:     /books
  Preset photos:    /photos
  Default:          /uploads

  The modified versions with transformations are stored under /t subdirectory.

  Usage
    $ cdn-upload <files>

  Options
    --preset, -p  photo preset to use
    --keep-name
`, {
  flags: {
    preset: {
      type: 'string',
      alias: 'p',
      default: 'default'
    },
    keepName: {
      type: 'boolean',
      default: false
    }
  }
})

function normalizeExtension (ext) {
  if (ext === '.jpg') return '.jpeg'
  return ext
}

function sha256 (data) {
  return createHash('sha256').update(data).digest('hex')
}

async function upload (data, filename) {
  if (!filename.startsWith('/')) {
    filename = '/' + filename
  }

  await put(`https://storage.bunnycdn.com/${config.zone}${filename}`, {
    headers: {
      AccessKey: config.key
    },
    body: data
  })

  return 'https://cdn.hacdias.com' + filename
}

; (async () => {
  const preset = presets[cli.flags.preset] || presets.default
  const path = presetsPath[cli.flags.preset] || presetsPath.default
  const files = cli.input

  if (cli.input.length === 0) {
    cli.showHelp()
    return
  }

  if (!presets[cli.flags.preset] && !presetsPath[cli.flags.preset]) {
    console.log('Preset', cli.flags.preset, 'does not exist.')
    process.exit(1)
  }

  for (const file of files) {
    console.log(file)
    const buff = await readFile(file)
    const ext = extname(file)
    const filename = basename(file, ext)

    if (!['.jpg', '.jpeg'].includes(ext) || cli.flags.preset === 'default') {
      const hash = cli.flags.keepName ? filename : sha256(buff)
      console.log(`\tUploading to ${path}/${hash}${normalizeExtension(ext)}`)
      console.log('\t', await upload(buff, `${path}/${hash}${normalizeExtension(ext)}`))
      continue
    }

    const img = sharp(buff)
    const metadata = await img.metadata()

    let jpeg = await img.jpeg({
      quality: 100,
      progressive: true
    })

    if (metadata.width > 3000) {
      jpeg = await jpeg.resize(3000)
    }

    const fBuff = await jpeg.toBuffer()

    console.log(`\tUploading to ${path}/${filename}${normalizeExtension(ext)}`)
    console.log('\t', await upload(fBuff, `${path}/${filename}${normalizeExtension(ext)}`))

    for (const type in preset) {
      for (const sizes of preset[type]) {
        const trans = sharp(buff)[type](typeOpts[type] || {}).resize(...sizes)
        const outName = `${filename}-${sizes.join('x')}${sizes.length === 1 ? 'x' : ''}.${type}`
        console.log(`\tUploading to ${path}/t/${outName}`)
        console.log('\t', await upload(trans, `${path}/t/${outName}`))
      }
    }
  }
})()
