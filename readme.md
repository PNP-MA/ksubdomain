ksubdomain is a stateless subdomain blasting tool, similar to stateless port scanning, which supports fast DNS blasting on Windows/Linux/Mac and has a retransmission mechanism so you don't have to worry about packet leakage.

The src asset collection of hacking8 information flow is https://i.hacking8.com/src/ and uses ksubdomain

![](image.gif)
## Install
1. Download the binary from https://github.com/boy-hack/ksubdomain/releases
2. Install libpcap environment
   - Windows
     Download the `npcap` driver. Some people report that the winpcap driver is invalid.
   - Linux
     libpcap has been statically compiled and packaged, no other operations are required
   - MacOS
     Comes with libpcap, no other operations required
3. Execute!
### Quick Installation
You need to have `go 1.17` or higher and install the `libpcap` environment. Run the following command
```
go install -v github.com/boy-hack/ksubdomain/cmd/ksubdomain@latest
```

##Useage
```bash
NAME:
   KSubdomain - Stateless subdomain brute force tool

USAGE:
   ksubdomain [global options] command [command options] [arguments...]

VERSION:
   1.8.6

COMMANDS:
   enum, e enumeration domain name
   verify, v verification mode
   test Test the maximum sending speed of the local network card
   help, h Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h show help (default: false)
   --version, -v print the version (default: false)

```

### model

**Verification Mode**
Provides a complete list of domain names, ksubdomain is responsible for quickly obtaining results

```bash
./ksubdomain verify -h

NAME:
   ksubdomain verify - verification mode

USAGE:
   ksubdomain verify [command options] [arguments...]

OPTIONS:
   --filename value, -f value Verify domain name file path
   --band value, -b value Broadband downlink speed, can be 5M, 5K, 5G (default: "2m")
   --resolvers value, -r value DNS server file path, one DNS address per line
   --output value, -o value Output file name
   --silent After use, the screen will only output the domain name (default: false)
   --retry value The number of retries. If it is -1, it will keep retrying (default: 3)
   --timeout value timeout (default: 6)
   --stdin accept input from stdin (default: false)
   --only-domain, --od only print domain name, not IP (default: false)
   --not-print, --np Do not print domain name results (default: false)
   --dns-type value DNS type 1 is a record 2 is ns record 5 is cname record 16 is txt (default: 1)
   --help, -h show help (default: false)
```

```
Reading from a file
./ksubdomain v -f dict.txt

Read from stdin
echo "www.hacking8.com"|./ksubdomain v --stdin

Read ns records
echo "hacking8.com" | ./ksubdomain v --stdin --dns-type 2
```

**Enumeration Mode**
Only provide the first-level domain name, specify the domain name dictionary or use the ksubdomain built-in dictionary to enumerate all second-level domain names

```bash
./ksubdomain enum -h

NAME:
   ksubdomain enum - enumerates domain names

USAGE:
   ksubdomain enum [command options] [arguments...]

OPTIONS:
   --band value, -b value Broadband downlink speed, can be 5M, 5K, 5G (default: "2m")
   --resolvers value, -r value DNS server file path, one DNS address per line
   --output value, -o value Output file name
   --silent After use, the screen will only output the domain name (default: false)
   --retry value The number of retries. If it is -1, it will keep retrying (default: 3)
   --timeout value timeout (default: 6)
   --stdin accept input from stdin (default: false)
   --only-domain, --od only print domain name, not IP (default: false)
   --not-print, --np Do not print domain name results (default: false)
   --dns-type value DNS type 1 is a record 2 is ns record 5 is cname record 16 is txt (default: 1)
   --domain value, -d value Domain name to blast
   --domainList value, --dl value Specify domain names from a file
   --filename value, -f value dictionary path
   --skip-wild Skip wild domain name resolution (default: false)
   --level value, -l value enumerates the domain name level, the default is 2, the second-level domain name (default: 2)
   --level-dict value, --ld value Enumerates dictionary files for multi-level domain names. Used when level is greater than 2. If not filled in, the default value will be used.
   --help, -h show help (default: false)
```

```
./ksubdomain e -d baidu.com

Get from stdin
echo "baidu.com"|./ksubdomain e --stdin
```

## Features and Tips

- No state blasting, with failure retransmission mechanism, extremely fast speed
- Chinese help, -h will show Chinese help
- Two modes, enumeration mode and verification mode, the enumeration mode has a built-in 10w dictionary
- Simplified the network parameters to -b parameters. Enter your network download speed, such as -b 5m, and the network card packet sending speed will be automatically limited.
- You can use ./ksubdomain test to test the maximum number of local packets sent
- Network card acquisition has been changed to fully automatic and can be read based on the configuration file.
- There will be a real-time progress bar, showing success/send/queue/receive/failure/time-consuming information in sequence.
- For data of different sizes, adjust the --retry --timeout parameters to achieve the best results
- When --retry is -1, it will keep retrying until all succeed.
- Support blasting ns records

## Comparison with massdns and dnsx

Using 1 million dictionaries, tested in a 4H5M network environment

| | ksubdomain | massdns | dnsx |
| -------- | ---------------------------------------- -------------------- | -------------------------------- ---------------------------------- | ------------------ ------------------------------------------------ |
| Supported systems | Windows/Linux/Darwin | Windows/Linux/Darwin | Windows/Linux/Darwin |
| Functionality | Supports validation and enumeration | Can only validate | Can only validate |
| Packet sending method | pcap network card packet sending | epoll,pcap,socket | socket |
| Command line | time ./ksubdomain v -b 5m -f d2.txt -o ksubdomain.txt -r dns.txt --retry 3 --np | time ./massdns -r dns.txt -t AAAA -w massdns .txt d2.txt --root -o L | time ./dnsx -a -o dnsx.txt -r dns.txt -l d2.txt -retry 3 -t 5000 |
| Notes | Added --np to prevent excessive printing | | |
| Result | Time taken: 1m28.273s<br />Successful: 1397 | Time taken: 3m29.337s<br />Successful: 1396 | Time taken: 5m26.780s<br />Successful: 1396 |

ksubdomain only takes 1 and a half minutes, which is much faster than massdns and dnsx~

## refer to

- Original ksubdomain https://github.com/knownsec/ksubdomain
- From Masscan, Zmap source code analysis to development practice <https://paper.seebug.org/1052/>
- Introduction to ksubdomain stateless domain name blasting tool <https://paper.seebug.org/1325/>
- [Comparison between ksubdomain and massdns](https://mp.weixin.qq.com/s?__biz=MzU2NzcwNTY3Mg==&mid=2247484471&idx=1&sn=322d5db2d11363cd2392d7bd29c679f1&chksm=fc986d10cbefe406f4bda22f62a16f08c71f31c241024fc82ecbb8e41c9c7188cfbd71276b81&token=76024279&lang=zh_CN#rd)
