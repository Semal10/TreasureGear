[1ns] -> CPU registers, CPU clock cycle
[1ns, 10ns] -> L1, L2 cache, expensive CPU operation, Branch misprediction penalty
[10ns, 100ns] -> L3 cache (low end), Apple M1 chip memory access (at the high end)
[100ns, 1000ns] -> sys call, MD5 hash
[1us, 10us] -> context switch btw linux thread, copying 64KB from one memory to other
[10us, 100us] -> Nginx, reading 1MB data sequential from memory, read from SSD
[100us, 1000us] -> SSD writes, intra zone round trip, Redis/Memcache get op
[1ms, 10ms] -> inter zone round trip (us-west-2a -> us-west-2b), HDD read
[10ms, 100ms] -> us-west to us-east trip, 1GB read sequential from memory
[100ms, 1000ms] -> encryption, TLS handshake, 1GB sequential read from SSD
[1s+] -> Data transfer of 1GB from one AZ to other in same region