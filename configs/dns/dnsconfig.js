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
  // Fastmail
  MX('@', 10, 'in1-smtp.messagingengine.com.'),
  MX('@', 20, 'in2-smtp.messagingengine.com.'),
  SRV('_submission._tcp', 0, 1, 587, 'smtp.fastmail.com.'),
  SRV('_imap._tcp', 0, 0, 0, '.'),
  SRV('_imaps._tcp', 0, 1, 993, 'imap.fastmail.com.'),
  SRV('_pop3._tcp', 0, 0, 0, '.'),
  SRV('_pop3s._tcp', 10, 1, 995, 'pop.fastmail.com.'),
  SRV('_jmap._tcp', 1, 1, 443, 'api.fastmail.com.'),
  SRV('_carddav._tcp', 0, 0, 0, '.'),
  SRV('_carddavs._tcp', 0, 1, 443, 'carddav.fastmail.com.'),
  SRV('_caldav._tcp', 0, 0, 0, '.'),
  SRV('_caldavs._tcp', 0, 1, 443, 'caldav.fastmail.com.'),
  CNAME('fm1._domainkey', 'fm1.hacdias.com.dkim.fmhosted.com.', CF_PROXY_OFF),
  CNAME('fm2._domainkey', 'fm2.hacdias.com.dkim.fmhosted.com.', CF_PROXY_OFF),
  CNAME('fm3._domainkey', 'fm3.hacdias.com.dkim.fmhosted.com.', CF_PROXY_OFF),
  TXT('@', 'v=spf1 include:spf.messagingengine.com ?all')
)

D('hacdia.sh', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  A('ipfs', THOR_IP, CF_PROXY_ON),
  A('xkcd', THOR_IP, CF_PROXY_ON),
  IGNORE_NAME('_dnslink.xkcd'),
  // Configuration for domains that do not send e-mails
  TXT('@', 'v=spf1 -all'),
  TXT('*._domainkey', 'v=DKIM1; p='),
  TXT('_dmarc', 'v=DMARC1; p=reject; sp=reject; adkim=s; aspf=s')
)

D('h4c.pt', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  // Configuration for domains that do not send e-mails
  TXT('@', 'v=spf1 -all'),
  TXT('*._domainkey', 'v=DKIM1; p='),
  TXT('_dmarc', 'v=DMARC1; p=reject; sp=reject; adkim=s; aspf=s')
)

D('henriquedias.com', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  // Configuration for domains that do not send e-mails
  TXT('@', 'v=spf1 -all'),
  TXT('*._domainkey', 'v=DKIM1; p='),
  TXT('_dmarc', 'v=DMARC1; p=reject; sp=reject; adkim=s; aspf=s')
)

D('hacdias.dev', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  // Configuration for domains that do not send e-mails
  TXT('@', 'v=spf1 -all'),
  TXT('*._domainkey', 'v=DKIM1; p='),
  TXT('_dmarc', 'v=DMARC1; p=reject; sp=reject; adkim=s; aspf=s')
)

D('nata.cafe', REG_NONE, DnsProvider(CLOUDFLARE),
  A('@', THOR_IP, CF_PROXY_ON),
  A('www', THOR_IP, CF_PROXY_ON),
  A('miniflux', THOR_IP, CF_PROXY_ON),
  // Configuration for domains that do not send e-mails
  TXT('@', 'v=spf1 -all'),
  TXT('*._domainkey', 'v=DKIM1; p='),
  TXT('_dmarc', 'v=DMARC1; p=reject; sp=reject; adkim=s; aspf=s')
)
