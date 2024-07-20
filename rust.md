# String
1. String read:
```rs
fn main() {
    // collect the characters as a vector of char
    let chars = "hi 🦀".bytes().collect::<Vec<u8>>(); // convert to ASCII
    for char in chars {
        print!("{} ", char); // 1 bytes => 8 bits
    }
    println!("");
    let tokens: Vec<char> = "hi 🦀".chars().collect(); // char is utf-8 by default
    for token in tokens {
        print!("{} ", token as u32); // 4 bytes => 32 bits
    }
}
```
Output:
```powershell
104 105 32 240 159 166 128 
104 105 32 129408 # the first three elements is the same
```
2. String as parameter
`String` and `&str` and `&'static str` are all utf-8 by efault, and can be passed to function through slice.
```rs
fn say_it_loud(msg:&str){
    println!("{:?}!!!",msg.to_string().to_uppercase().chars());
}

fn main() {
    // say_it_loud can borrow &'static str as a &str
    say_it_loud("hello");
    // say_it_loud can also borrow String as a &str
    say_it_loud(&String::from("ヽ(* ﾟ д ﾟ)ノ"));
}
```
3. String builtin functions
* `concat()`: concat an array to a &str, eg. `helloworld = ["hello", " ", "world", "!"].concat();`
* `join()`: you can add something in `()` as an argument, eg. `abc = ["a", "b", "c"].join(",");`

# Debug
1. Type-name of a variable
```rs
use std::any::type_name;

// A function that prints the type name of its generic parameter
fn print_type_of<T>(_: T) {
    println!("The type is: {}", type_name::<T>());
}

fn main() {
    let my_string = String::from("  Hello, World!  ");
    
    // Using trim() and converting back to String
    let trimmed = my_string.trim();
    print_type_of(trimmed);
    print_type_of(my_string);
    
}
```

# OOP
1. Rust supports **polymorphism** with `traits`.
```rs
// Define a trait
trait Printable {
    fn print_info(&self);
}

// Define a function shows polymorphism
fn display_info<T: Printable>(item: T) {
    item.print_info();
}

// It doesn't care about the data type
display_info(person); // 输出: Person: Name - Alice, Age - 30
display_info(book);   // 输出: Book: Title - The Great Gatsby, Author - F. Scott Fitzgerald
```
More:
1. One can use `impl` to imply a method, and you also need a `&self`, which refer to the type **immediately** after `impl`
e.g. You can only use the function of the `trait`, and we need imply `make_noise()` for `SeaCreature` to get the data inside it.
```rs
trait NoiseMaker {
    fn make_noise(&self);
    
    fn make_alot_of_noise(&self){
        self.make_noise();
        self.make_noise();
        self.make_noise();
    }
}

impl NoiseMaker for SeaCreature {
    fn make_noise(&self) {
        println!("{}", &self.make_noise());
    }
}
``` 
# Graph
> Anology to cpp
```cpp
void add(int a, int b, int w) {
    e[cur] = b, ne[cur] = h[a], h[a] = cur++;
}
```
A BFS example is [here](https://codeforces.com/contest/1307/submission/271382297)
> More can be referred [here](https://github.com/EbTech/rust-algorithms/blob/master/README.md#graphs)

# Common I/O
1. Template in ICPC
```rs
// Author: ヽ(* ﾟ д ﾟ)ノ
#![allow(unused_imports)]
use std::cmp::{min,max};
use std::io::{self, Write};
use std::str;
use std::collections::HashSet;
use io::BufRead;

fn solve<R: BufRead, W: Write>(scan: &mut Scanner<R>, out: &mut W) {
    // code here
}

fn main() {
    let (stdin, stdout) = (io::stdin(), io::stdout());
    let mut scan = Scanner::new(stdin.lock());
    let mut out = io::BufWriter::new(stdout.lock());
    let T: usize = 1;
    // let T = scan.token::<usize>();
    for _ in 0..T{
        solve(&mut scan, &mut out);
    }
}

struct Scanner<R> {
    reader: R,
    buf_str: Vec<u8>,
    buf_iter: str::SplitAsciiWhitespace<'static>,
}
impl<R: io::BufRead> Scanner<R> {
    fn new(reader: R) -> Self {
        Self { reader, buf_str: Vec::new(), buf_iter: "".split_ascii_whitespace() }
    }
    fn token<T: str::FromStr>(&mut self) -> T {
        loop {
            if let Some(token) = self.buf_iter.next() {
                return token.parse().ok().expect("Failed parse");
            }
            self.buf_str.clear();
            self.reader.read_until(b'\n', &mut self.buf_str).expect("Failed read");
            self.buf_iter = unsafe {
                let slice = str::from_utf8_unchecked(&self.buf_str);
                std::mem::transmute(slice.split_ascii_whitespace())
            }
        }
    }
}
```
1. `vector`:
```rs
let pile: Vec<_> = (0..n).map(|_| scan.token::<usize>()).collect();
let mut t: Vec<i32> = vec![];
```
2. `String` as ASCII and other [tricks](https://codeforces.com/contest/1194/submission/269517484)
```rs
let p: Vec<u8> = scan.token::<String>().bytes().collect();
```