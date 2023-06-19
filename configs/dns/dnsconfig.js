var REG_NONE = NewRegistrar('none')
var CLOUDFLARE = NewDnsProvider('cloudflare')

var THOR_IP = '135.181.87.57'

D('hacdias.com', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  A('git', THOR_IP, CF_PROXY_ON),
  A('go', THOR_IP, CF_PROXY_ON),
  CNAME('cdn', 'hacdias.b-cdn.net.', CF_PROXY_OFF),
  TXT('@', 'google-site-verification=lg_VxNseU4oOVkGPbTQjGa0oHQbNyq-gI8Xf_gF5IA8'),
  TXT('_atproto', 'did=did:plc:xsx3bphrwkgeo3qnfjhzmdra'),
  // Migadu
  SRV('_autodiscover._tcp', 0, 1, 443, 'autodiscover.migadu.com.'),
  SRV('_submissions._tcp', 0, 1, 465, 'smtp.migadu.com.'),
  SRV('_imaps._tcp', 0, 1, 993, 'imap.migadu.com.'),
  SRV('_pop3s._tcp', 0, 1, 995, 'pop.migadu.com.'),
  CNAME('autoconfig', 'autoconfig.migadu.com.', CF_PROXY_OFF),
  CNAME('key1._domainkey', 'key1.hacdias.com._domainkey.migadu.com.', CF_PROXY_OFF),
  CNAME('key2._domainkey', 'key2.hacdias.com._domainkey.migadu.com.', CF_PROXY_OFF),
  CNAME('key3._domainkey', 'key3.hacdias.com._domainkey.migadu.com.', CF_PROXY_OFF),
  MX('@', 10, 'aspmx1.migadu.com.'),
  MX('@', 20, 'aspmx2.migadu.com.'),
  TXT('@', 'hosted-email-verify=93ksh76w'),
  TXT('@', 'v=spf1 include:spf.migadu.com -all'),
  TXT('_dmarc', 'v=DMARC1; p=quarantine;')
)

D('hacdia.sh', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  A('ipfs', THOR_IP, CF_PROXY_ON),
  A('xkcd', THOR_IP, CF_PROXY_ON),
  IGNORE_NAME('_dnslink.xkcd')
)

D('h4c.pt', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON)
)

D('henriquedias.com', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON)
)

D('nata.cafe', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  A('miniflux', THOR_IP, CF_PROXY_ON),
  A('linkding', THOR_IP, CF_PROXY_ON)
)
