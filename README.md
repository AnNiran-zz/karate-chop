# Brief description:
Karate-chop contains binary search logic implementations according to
technical specifications of 
http://codekata.com/kata/kata02-karate-chop/

Set up values for data length generation and highest random value - for setting up the range
are inside /binarysearch/config.go file 

# Running:
make

to run the program with finding the target key from generated data:
./cmd binary-search -algorithm <algorithm-name> -target <target-value>
select algorithm from the following (find description for each below):
    "flip-step"
    "pointers-move"
    "pointers-move-int"
    "recursive-dynamic"
    "iterate-dynamic"
    "parallel-dynamic"

Set up values for data length generation and highest random value - for setting up the range
are inside /binarysearch/config.go file 

*** Run ***
to run help display information:
./cmd binary-search -help

to output description for each implementation:
./cmd binary-search -info <algorithm-name>


*** Please find below description of each logic ***

=== Flip-step ===
"flip-step" implementation of binary search logic uses a pointer and a step with changing size to
search through the list
At the begining the pointer is set at the middle list position and the step size is set to half length of the list

-------- step size -------
- - - - - - - - - - - - - - - - - - - - - - - - -
						 |
						 p = step

After comparing the target with the value at position p of the list:
* if the target is smaller than the value at p:
the step size is cut with 50% if it is larger than 1
step is "flipped" to the smaller range of values in the list
            ---step size--
- - - - - - - - - - - - - - - - - - - - - - - - -
			|            |
			p1          <-

p is moved at the new position - decreasing its key with the new step size

* if the target is greater than the value at p:
step size is cut with 50% if it is larger than 1
the step is "flipped" to the range with larger values in the list
                   ---ss--
- - - - - - - - - - - - - - - - - - - - - - - - - 
            |      |     |
			p1     |    <-
			->     p2

At each step the step size is decreased with 50% which decreases the potential range of values
to be searched


=== Moving-pointers ===
"pointers-move" implementation for binary search algorithm searches through the sorted 
array of values for the given target using three pointers:
starting, ending and searching pointer - that are used to move accross the list after each binary check 

At the first step:
* starting pointer is set at the first position of the list, ending pointer - 
at the last position and the searching pointer is set at the position at the half length of the list

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
|                                   |                                  |
ps                                  p                                  pe

After comparing the list value at the searching pointer position with the target:
* if the target is greater - the searching pointer is moved towards the right with 
half of its previous value
* starting pointer is moved at the previous searching pointer position
* the ending pointer keeps its current position

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
|                                   |                                    |
ps ->                               p ->                                 pe
									
									ps                 p                 pe
									   
* if the target is smaller than the searching pointer positioin
* the searching pointer is moved towards the smallest range with half of its previous value 
* the ending pointer is moved at the previous position of the searching pointer
* the starting pointer keeps its position

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
|                                   |                                    |
ps ->                               p ->                                 pe
									
									ps                 p              <- pe

									ps     <-p      <-pe


At each loop the searching pointer key is cut with 50% and the range between starting and ending pointers decreases

If no value is found -1 is returned




=== Moving Pointers - interface ===
"pointers-move-int" implementation with interface functions in the same way 
as the "pointers-move", except that is uses custom defined types - 
a pointer for values of the three pointers, and a in interface
that the type implements

this is done only for prepresentational purposes and is not considerably important
for the code performance and resources optimisation




=== Recursive-dynamic ===
"recursive-dynamic" logic implements binary search algorithm by using 
separate data lists (slices) for keys of the originally generated data, as well as 
a copy of the data and a pointer

Initially the copy of the generated data contains all of it, as well as the keys slice 
contains all keys, the pointer is set at the half of the original data lenght position

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  keys
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  copy data
										|
										p

At each step when the value at the pointer position is compared to the target
If the target is greater than the value at p:
* the pointer is moved to towards the greater values
* length of the data and keys lists are cut with 50% of their lengths

                                        - - - - - - - - - - - - - - - - - - - - - - -  keys
                                        - - - - - - - - - - - - - - - - - - - - - - -  copy data
										|
										p ->                      |
																  p
											
If the target is smaller than the value at p:
* the pointer is moved towards the smaller values with half of its previous size
* data copy and keys parts after the previous pointer position are removed

                                        - - - - - - - - - - - - - -                    keys
                                        - - - - - - - - - - - - - -                    copy data
                                        |
                                        p ->                      |
						                            |          <- p
													p
													
After each step the size of keys and data copy decreases as well as the potential range values
for searcging
When the target is met - the value at the p from keys list is returned, otherwise -1




=== Dynamic iteration ===
"iterate-dynamic" implementation of the binary search algoritm uses minimal values settings:
an offset for calculating next range for iteration

At first the offset value is set at 0
If the target is greater than the value at p postion:
* the remaining data with smaller values is cut from the list
* offset is moved to the larger values with half of its previous size

If the target is smaller than the value at p:
* the data part before the last position - containing smaller values, is removed
* offset is moved to the larger values direction with half of its size

sData is dynamically cut at each step while its length decreases => that is how the potential range is smaller each time
and target key is found



=== Parallel Dynamic ===
"parallel-dynamic" implementation of the binary search algorithm uses:
* a map of two sub-slices of the data to make the checks at each step
sub-slices are copy of the original generated list (slice) and each of them contains half of the data
* a map of two sub-slices of data value keys from the original list; each slice in the map is half
of the size of the original data at the beginning, and matches the values in the oher map
with their original keys
* a pointer to move across the two subsices - rails of the data copy and compare values at their positions
with the target

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - original data slice

- - - - - - - - - - - - - - - - - - - - - - - -
- - - - - - - - - - - - - - - - - - - - - - - - map of two slices with keys matching the values inside the map with data copy

- - - - - - - - - - - - - - - - - - - - - - - -
- - - - - - - - - - - - - - - - - - - - - - - - map of two slices, copy of original data
					   |
					   p

At each step continuously updates each "rail" of the map containing the data,
as well as the "rails" inside the map with the keys
If the target is greater than both values at p position inside both rails with data:
* the data "rails" with greater values are kept - rest are moved
* keys "rails" are also updated to contain the keys for the values inside the data map slices
* pointer is moved to the right with half of its previous size and points to the new value
at the laf of the data map slices

                        - - - - - - - - - - - -
                        - - - - - - - - - - - - map of two slices with keys matching the values inside the map with data copy

                        - - - - - - - - - - - -
                        - - - - - - - - - - - - map of two slices, copy of original data
					   |
					   p -> 
					                p

If the target is smaller than both values at p position inside both rails:
* the data "rails" with smaller values are kept - rest are removed
* keys "rails" are also updated to contain the keys for the values inside the data map slices
* pointer is moved to the right with half of its previous size and points to the new value
at the laf of the data map slices

- - - - - - - - - - - -
- - - - - - - - - - - -                        map of two slices with keys matching the values inside the map with data copy

- - - - - - - - - - - -
- - - - - - - - - - - -                        map of two slices, copy of original data
                       | 
                    <- p
		   p
		   
If the target is smaller than value at p position in one of the "rails" and greater than the other:
* the corresponding halves of each "rail" are taken and the data map is updated to 
contain the new data
* keys map is updated as well to mirror the data map - to hold the keys of the values from the original slice
* pointer is moved with half of its previous value
 
                        - - - - - - - - - - - -
- - - - - - - - - - - -                         map of two slices with keys matching the values inside the map with data copy

                        - - - - - - - - - - - -
- - - - - - - - - - - -                         map of two slices, copy of original data
                       | 
                    <- p
		   p

- - - - - - - - - - - -                        
- - - - - - - - - - - -                         map of two slices with keys matching the values inside the map with data copy

- - - - - - - - - - - -
- - - - - - - - - - - -                         map of two slices, copy of original data
		  |
		  p


After that the check of p with the values at p is done again and data is cut in halves each time
When the target is met - its key is extracted from the keys map is returned, otherwise -1

