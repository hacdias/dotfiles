// If you're asking if I should have this public, I don't see why not.
// All DNS records are public.
var REG_NONE = NewRegistrar('none', 'NONE')
var CLOUDFLARE = NewDnsProvider('cloudflare', 'CLOUDFLAREAPI')

var THOR_IP = '135.181.87.57'

D('hacdias.com', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('miniflux', THOR_IP, CF_PROXY_ON),
  A('xkcd', THOR_IP, CF_PROXY_ON),
  A('ownyourtrakt', THOR_IP, CF_PROXY_ON),
  A('git', THOR_IP, CF_PROXY_ON),
  CNAME('www', 'hacdias.com.', CF_PROXY_ON),
  CNAME('cdn', 'hacdias.b-cdn.net.', CF_PROXY_OFF),
  CNAME('key1._domainkey', 'key1.hacdias.com._domainkey.migadu.com.', CF_PROXY_OFF),
  CNAME('key2._domainkey', 'key2.hacdias.com._domainkey.migadu.com.', CF_PROXY_OFF),
  CNAME('key3._domainkey', 'key3.hacdias.com._domainkey.migadu.com.', CF_PROXY_OFF),
  MX('@', 10, 'aspmx1.migadu.com.'),
  MX('@', 20, 'aspmx2.migadu.com.'),
  TXT('@', 'hosted-email-verify=93ksh76w'),
  TXT('@', 'keybase-site-verification=3OP1gzrj1B90Z-TuNZe5piO-5jp3UP7OCAEEG7yiPWE'),
  TXT('@', 'google-site-verification=lg_VxNseU4oOVkGPbTQjGa0oHQbNyq-gI8Xf_gF5IA8'),
  TXT('@', 'v=spf1 include:spf.migadu.com -all'),
  TXT('krs._domainkey', 'k=rsa; p=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC9ZfGRGj9sn0GnSq6lZTZyfgkujmDdBkLEXgeJ8xbloPSX6EHUfieHkrTW9wpFm8JGVC1SONTIrs3NDuyaIR/c62IicylIQIlFqJ0hKnfXmRq1SEu6m7WvbnaJ12x605oC4RP0fa0/do7QOnpgAQDnZPDoRuRqc+p2OJQPY8+YtwIDAQAB'),
  TXT('_dmarc', 'v=DMARC1; p=quarantine;'),
  TXT('_dnslink.xkcd', 'dnslink=/ipfs/bafybeigyzt2b5olvd6upxkro26fechbn5hsxmnt5i2jchmn76vmxpfdjde')
)

D('hacdia.sh', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON)
)

D('h4c.pt', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON)
)

D('henriquedias.com', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  CNAME('www', 'henriquedias.com.', CF_PROXY_ON),
  TXT('_keybase', 'keybase-site-verification=0ajL5eM3gy5vO10gNlmGqhFvVuxvqOuRYDQIFWFCDaw'),
  TXT('@', 'google-site-verification=jIEoWFXm813Y2EAD4GzvNHrf1uOWJmsWZlsF9E2UgsM')
)

D('nata.cafe', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON)
)
