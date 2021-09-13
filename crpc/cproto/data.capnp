@0x8bc6a44b4908aa13;
using Go = import "/go.capnp";

$Go.package("cpb");
$Go.import("benchprotos/crpc/cproto");

struct Quote @0x8623644ff5101c64 {  # 8 bytes, 2 ptrs
  ack @0 :Text;  # ptr[0]
  bid @1 :Text;  # ptr[1]
  price @2 :Float32;  # bits[0, 32)
}
interface QuotesUpdate @0xecf32a66a83572cc {
  quotesUpdate @0 (list :List(Quote)) -> (result :Text);
}


