package main

import (
    "bufio"
    "fmt"
    "os"
)

/* Global slice of maps containing parse results */

var results []map[string]int
 

/* Function for processing an individual file to generate map of word frequency 
 * Inputs: (1) File name  (2) Index within list of files (always called with 0 in this version of the program) 
*/

func process_file(file_name string, i int) {

    // Open the input file.

    fi, err := os.Open(file_name)

    if err != nil {
        fmt.Printf("Input file not found!\n")
        return
    }
    defer fi.Close()

    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(fi)

    // Split by words using built-in word scanner in bufio package 
    // Could have used alternate word scan techniques as well, this one uses unicode white space definitions

    scanner.Split(bufio.ScanWords)

    // Now make the map results */

    results[i] = make(map[string]int)

    for scanner.Scan() {
        word := scanner.Text()
        results[i][word]++
//        fmt.Println(word)
    }


}


func main() {

    fi_name := "./input_files/moby-000.txt"

    // Open the output file.
    fo, err := os.Create("exercise1-output.txt")

    if err != nil {
        fmt.Printf("Error creating output file!\n")
        return
    }
    defer fo.Close()

    /* Since we know we will only deal with 1 file and 1 resultant map in this version of the program, 
     * lets allocate a slice of 1 map 
    */ 

    results = make([]map[string]int, 1)

    /* Process the input file */
    i := 0 
    process_file(fi_name, i) 

    for word, word_count := range results[i] {
        fmt.Printf("%v %v \n", word, word_count)
        fmt.Fprintf(fo, "%v %v \n", word, word_count)
    }

}

