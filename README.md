# nest
nest is a Go library to work on deeply nested datastructures in a more managed way.

  nest is a library that provides an easy way to retreive fields from deeply nested
  data structures. Structure here can be of any Go type that can store data (slice,
  maps, structs, primitives).
  
  Fields are accessed by using a specific path string. The path string contains
  accessors for the nested structures seperated by `/`.
  
  Accessors can be of two types:
    1. Direct field/index accessor.
    2. Special types.
    
  Each path string must begin with a `/`. This `/` is the root structure. Simply getting
  the path `/` returns the original structure. e.g.:
  
    var dst = []int{}
    var src = []int{1, 2, 3}
    Get("/", src, &dst) // dst == []int{1, 2, 3} == src -- deep copy
  
  
  1. Direct field/index accessors:
  After the root path, direct accessor can be used. It can be an index, in case of a slice,
  or a field name/map key. e.g.:
  
    var dst int
    Get("/1", src, &dst) // dst == 2 == src[1]
  
    type Course struct {
      Name      string
      Teacher   string
    }
    var src = struct{
      Name    string
      Age     int
      Marks   []int
      Courses []Course
    }{
      Name:  "Abh",
      Age:   10,
      Marks: []int{20, 19, 15}
      Courses: []Course{ {"Physics", "Prof. P"} {"Chemistry", "Prof. C"} }
    }
  
    var dst string
    Get("/Name", src, &dst) // dst == "Abh" == src.Name
  
    var dst int
    Get("/Marks/2", src, &dst) // dst == 15 == src.Marks[2]
  
    var dst string
    Get("/Courses/0/Teacher", src, &dst) // dst == "Prof. P"
  
  
2. Special Type Accessors: Use of only direct accessor result in single values,
   as illustrated above. To get multiple values, special accessors are required.
   There are two types of these accessors:
     
     
  **The Dot (`.`) Accessor:**
  The Dot accessor fetches each element in an iterable type (slice, map). For example,
  consider the following value:
     
    type InnerSimple struct {
      I int
      S string
    }
     
    v := []InnerSimple{{1, "one"}, {2, "two"}, {3, "three"}}
     
Now, to get all the values in v, we can write:
     
    var result []InnerSimple
    Get("/.", v, &result) // result == v
     
But this is not a very useful result. But we can see that '.' fetched each value in v
and put it in a new slice. Now consider the case when you want all the string values
from the slice:
     
    var result []int
    Get("/./S", v, &result) // result == []string{"one", "two", "three"}
     
Now this is a more useful result. `.` can be used multiple times in a path as long as
it follows an iterable type in the path. In above path, `/` is a slice itself and is
iterable. The `.` accesses each element in that slice (`/` here) and fetches the elements
'S' field. Note that each element of the slice must either be a struct with a field named
'S' or a map with a key 'S'.
     
As another example, consider:
     
    type Simple struct {
      SimpleI int
      SimpleS string
      InnerSimple
      SimpleSlc []InnerSimple
    }
    v := []Simple{
      {1, "one", InnerSimple{10, "ten"}, []InnerSimple{InnerSimple{1, "a"}, InnerSimple{2, "b"}}},
      {2, "two", InnerSimple{20, "twenty"}, []InnerSimple{InnerSimple{3, "c"}, InnerSimple{4, "d"}}},
      {3, "three", InnerSimple{30, "thirty"}, []InnerSimple{InnerSimple{5, "e"}, InnerSimple{6, "f"}}},
    }
     
    var result [][]string
    Get("/./SimpleSlc/./S", v, &result) // result ==  [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
     
In above example, the path starts at root, which is a slice (`[]Simple`). It then accesses
each element of this slice. For each element, it accesses the SimpleSlc field. This
field is also a slice and so another `.` can be used to access each element of SimpleSlc.
After the second `.`, 'S' field of each element of SimpleSlc is accessed.
 
Note the type of the result. It is of `[][]string` type. That's because each `.` fills up
exactly one slice. Inner `.` works on a single SimpleSlc and returns a slice containing
all the 'S' values. The outer `.` collect these values in another slice. Hence `[][]string`.
But what if we wanted to merge the inner slices together. That's where our second
accessor type comes into play.
    
     
**The Star (`*`) Accessor:**
There are two important rules associated with this accessor.
a) Star accessor can only be used to replace a '.' in the path.
b) Star accessor cannot be the last accessor (or only) in a path. I.e., path '/A/B/* /C'
   is illegal.
What '*' does is that it breaks the structure or returned result. If in above path
`/./SimpleSlc/./S`, we replace the first '.' with a star, we get the following result:
     
    var result [][]string
    Get("/* /SimpleSlc/./S", v, &result) // result == []string{"a", "b", "c", "d", "e", "f"}
     
As evident, '*' merges together the outer slice. This is why '*' must always be used
before a '.', because it must have some result to merge together.
Many more examples can be found in the test files, which also include map accesses.
     
  
TODO:<br>
Add data updates through path.<br>
Add custom function to process accessed data.<br>
Add interface similar to Marshalling/Unmarshalling similar to encoding.<br>
