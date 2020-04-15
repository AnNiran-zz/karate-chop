/*
binarysearch package contains multiple implementations of binary searching
Follows up the specifications in Code Kata 02 at:
http://codekata.com/kata/kata02-karate-chop/

Set up values for data length generation and highest random value - for setting up the range
are inside /binarysearch/config.go file 

*** Please find below description of each logic ***

=== Moving-pointers ===
moving-pointers is a logic for binary search algorithm implementation
that searches through a sorted array of values for the given target
The logic uses 3 pointers - starting, ending and searching pointer - that are used to move
accross the list after each binary check 

At the first step starting pointer is set at the first position of the list, ending pointer - 
at the last position and the searching pointer is set at the position at the half length of the list

After comparing the list value at the searching pointer position with the target:
* if the target is greater - the searching pointer is moved towards the right with 
half of its previous value, starting pointer is moved at the previous searching pointer position
and the ending pointer keeps its current position

* if the target is smaller than the searching pointer positioin - it is moved towards the left 
with half of its previous value and the ending pointer is moved at the previous position of the 
searching pointer, the starting pointer keeps its position

At each loop the check is made, as well as if the value at the searching pointer position is 
matching the target
If no value is found -1 is returned


=== Moving Pointers - interface ===
moving-pointers implementation with interface functions in the 
same way as the "moving-pointers", except that is uses custom defined
types - pointer for values of the three pointers, and a in interface
that the type implements
this is done only for prepresentational purposes and is not considerably important
for the code performance and resources optimisation


=== Flip Step ===
Flip-step implementation of binary search logic uses a pointer and a stpe with changing size to
search through the list 
At the begining the pointer is set at the middles of the list position and the step size
is set to half length of the list

After comparing the target with the value at position p of the list:
* if the target is smaller than the value at p - the step size is "flipped" to the smaller range
of values in the list
* if the target is greater than the value at p - the step is "flipped" to the range with larger values

at each step the step size is decreased with 50% which decreases the potential range of values


=== Recursive-dynamic ===
recursive dynamic logic implements binary search algorithm by using 
separate data lists (slices) for keys of the originally generated data, as well as 
a copy of the data and a pointer

Initially the copy of the generated data contains all of it, as well as the keys slice 
contains all keys, the pointer is set at the half of the original data lenght position

At each step when the value at the pointer position is compared to the target - 
the pointer is moved tho the right if the target is greater than the value, and to the left -
if the target is smaller than the value respectively
After that the remaining part of the copied data in the opposite direction is removed, as well
as its corresponding keys from the second list 

When a match is found the corresponding key is returned, if not - -1


=== Dynamic iteration ===
"iterate-dynamic" implementation of the binary search algoritm uses minimal values settings:
an offset for calculating next range for iteration

At first the offset value is set at 0 and after each iteration, the sData content that is not to be checked
is removed and the offset is updated:
* if the value at the p (pointer) position is smaller than the target - offset is updated to cover half of the 
data plus one position for the remaining part of it
* if the target value is smaller than the value at the p position - the offset keeps its size

sData is dynamically cut at each step while its length decreases => that is how the potential range is smaller each time
and target key is found 


=== Parallel Dynamic ===
parallel-dynamic' implementation of the binary search algorithm uses a map of two sub-slices of the data
to make the checks at each step
sub-slices are copy of the original generated list (slice), as well as two sub-slices of their keys are kept inside 
a separate map

uses a pointer to make the comparison with the target at the pointer value of the lists in the same manner,
but instead of cutting one list, uses two and continuously updates each "rail" of the map containing the data,
as well as the "rails" inside the map with the keys

* if the target is greater than both values at p position inside both rails with data - the parts holding 
greater values are kept for the next step
* if the target is smaller than the both values at p position inside both rails with data - the parts holding 
smaller values are kept for the nest step
* if target is smaller than one of the values at p position in the rails, and greater than the other - 
respective part is kept for the next step - from one rail - the part that holds larger values and from the other -
the one holding smaller ones

At each step the keys data map is updated in the same manner
When the target is met - its key is extracted from the keys map and returned, otherwise -1



*/
package binarysearch
