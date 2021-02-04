#!/usr/bin/env node
'use strict'

const home = require('os').homedir()
require('dotenv').config({ path: home + '/scripts/.env' })

const got = require('got')
const meow = require('meow')
const crypto = require('crypto')
const { extname } = require('path')
const fs = require('fs-extra')
const sharp = require('sharp')
const { basename } = require('path')

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
  books: {
    jpeg: [
      [100, 150],
      [200, 300]
    ],
    webp: [
      [100, 150],
      [200, 300]
    ]
  },
  default: {
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
  books: 'books',
  photos: 'photos',
  default: 'uploads'
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
`, {
  flags: {
    preset: {
      type: 'string',
      alias: 'p',
      default: 'default'
    }
  }
})

function normalizeExtension (ext) {
  if (ext === '.jpg') return '.jpeg'
  return ext
}

;(async () => {
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
    const buff = await fs.readFile(file)
    const ext = extname(file)
    const filename = basename(file, ext)

    if (!['.jpg', '.jpeg'].includes(ext)) {
      const hash = sha256(buff)
      console.log(`\tUploading to ${path}/${hash}${normalizeExtension(ext)}`)
      console.log('\t', await upload(buff, `${path}/${hash}${normalizeExtension(ext)}`))
      continue
    }

    const fBuff = await sharp(buff).jpeg({
      quality: 100,
      progressive: true
    }).resize(3000).toBuffer()

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

function sha256 (data) {
  return crypto.createHash('sha256').update(data).digest('hex')
}

async function upload (data, filename) {
  if (!filename.startsWith('/')) {
    filename = '/' + filename
  }

  await got.put(`https://storage.bunnycdn.com/${config.zone}${filename}`, {
    headers: {
      AccessKey: config.key
    },
    body: data
  })

  return 'https://cdn.hacdias.com' + filename
}