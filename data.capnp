 @0xf1685067cedef0e0;
$package("main");
$import("github.com/jkrus/benchprotos");

struct Date {
  # A standard Gregorian calendar date.

  year @0 :Int16;
  # The year.  Must include the century.
  # Negative value indicates BC.

  month @1 :UInt8;   # Month number, 1-12.
  day @2 :UInt8;     # Day number, 1-30.
}