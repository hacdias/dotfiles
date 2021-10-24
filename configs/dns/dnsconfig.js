// If you're asking if I should have this public, I don't see why not.
// All DNS records are public.
var REG_NONE = NewRegistrar('none', 'NONE')
var CLOUDFLARE = NewDnsProvider('cloudflare', 'CLOUDFLAREAPI')

var THOR_IP = '135.181.87.57'

D('hacdias.com', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('miniflux', THOR_IP, CF_PROXY_ON),
  A('xkcd', THOR_IP, CF_PROXY_ON),
  CNAME('www', 'hacdias.com.', CF_PROXY_ON),
  CNAME('status', 'stats.uptimerobot.com.', CF_PROXY_ON),
  CNAME('cdn', 'hacdias.b-cdn.net.', CF_PROXY_OFF),
  CNAME('email', 'mailgun.org.', CF_PROXY_OFF),
  CNAME('dkimprovmx1._domainkey', 'dkimprovmx1.improvmx.com.', CF_PROXY_OFF),
  CNAME('dkimprovmx2._domainkey', 'dkimprovmx2.improvmx.com.', CF_PROXY_OFF),
  CNAME('drive', 'nx14546.your-storageshare.de.', CF_PROXY_OFF),
  MX('@', 10, 'mx1.improvmx.com.'),
  MX('@', 20, 'mx2.improvmx.com.'),
  TXT('@', 'keybase-site-verification=3OP1gzrj1B90Z-TuNZe5piO-5jp3UP7OCAEEG7yiPWE'),
  TXT('@', 'google-site-verification=lg_VxNseU4oOVkGPbTQjGa0oHQbNyq-gI8Xf_gF5IA8'),
  TXT('@', 'v=spf1 include:mailgun.org include:spf.improvmx.com ~all'),
  TXT('krs._domainkey', 'k=rsa; p=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC9ZfGRGj9sn0GnSq6lZTZyfgkujmDdBkLEXgeJ8xbloPSX6EHUfieHkrTW9wpFm8JGVC1SONTIrs3NDuyaIR/c62IicylIQIlFqJ0hKnfXmRq1SEu6m7WvbnaJ12x605oC4RP0fa0/do7QOnpgAQDnZPDoRuRqc+p2OJQPY8+YtwIDAQAB'),
  TXT('_dmarc', 'v=DMARC1; p=none;'),
  TXT('_dnslink', 'dnslink=/ipfs/QmWADWrbLF7gPDxBNacj4K5cSBZvz2tb7G7uKEvgMWduZ1'),
  TXT('_dnslink.xkcd', 'dnslink=/ipfs/QmV4V1xkgosbHBL1eq2p4GSr95174Y41H88gaoqQX5L6xo')
)

D('hacdia.sh', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON)
)

D('henriquedias.com', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  CNAME('www', 'henriquedias.com.', CF_PROXY_ON),
  MX('@', 10, 'mx1.improvmx.com.'),
  MX('@', 20, 'mx2.improvmx.com.'),
  TXT('_keybase', 'keybase-site-verification=0ajL5eM3gy5vO10gNlmGqhFvVuxvqOuRYDQIFWFCDaw'),
  TXT('@', 'google-site-verification=jIEoWFXm813Y2EAD4GzvNHrf1uOWJmsWZlsF9E2UgsM'),
  TXT('@', 'v=spf1 include:spf.improvmx.com ~all')
)
