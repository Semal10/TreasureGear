DNS is internet directory

Sole purpose of translating Domain Name -> IP Address


What happens when we search google.com?

Browser Cache -> OS Cache -> DNS Resolver (ISP/Cloudflare(1.1.1.1)/Google(8.8.8.8)) [It has also it's cache] -> Root Nameserver -> TLD Nameservers (.com/.org/.edu etc) -> Authoritative Name Server -> IP

Gotchas
1) Shorten the TTL before live update to high traffic DNS
2) Keep the old IP machine running for a while as DNS propogation takes time