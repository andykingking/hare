@0x9aa94e0b6599d365;
using Go = import "go.capnp";
$Go.package("hare");
$Go.import("testpkg");


struct DocumentCapn { 
   body  @0:   Text; 
} 

struct SequencerCapn { 
   mark  @0:   UInt64; 
} 

struct ViewCapn { 
   src  @0:   Text; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

