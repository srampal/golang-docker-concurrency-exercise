package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sync"
)

/* Global slice of maps containing parse results. Each element of the slice is a map for a unique file */

var results []map[string]int
 
/* Overall/ master results map. Populated by aggregating data from each of the results[i] maps */
var master_results map[string]int

/* Function for processing an individual file to generate map of word frequency 
 * Inputs: (1) File name  (2) Index within list of files 
*/

func process_file(file_name string, i int, wg *sync.WaitGroup) {

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

    wg.Done()
}


func main() {


    // Open the output file.
    fo, err := os.Create("exercise2-output.txt")

    if err != nil {
        fmt.Printf("Error creating output file!\n")
        return
    }
    defer fo.Close()

    /* Read all files in this directory and make a list of their names */

    dir, err := os.Open("./input_files")
    if err != nil {
        fmt.Printf("Error opening input files directory!! \n")
        return
    }
    defer dir.Close()

    files_info, err := dir.Readdir(-1)
    if err != nil {
        fmt.Printf("Error reading input files directory!! \n")
        return
    }

    num_files := len(files_info)
    fmt.Printf(" Total files %d \n", num_files)

    dir_name := "./input_files"

    /* Now create a slice of num_files elements, where each element is the map for one of the files */ 
    results = make([]map[string]int, num_files)

    /* Create the master_results global map which will accumulate across all maps results[i] */ 
    master_results = make(map[string]int)

    var wg sync.WaitGroup

    wg.Add(num_files)

    for i, fi := range files_info {
        full_file_name := strings.Join([]string{dir_name, fi.Name()}, "/")
        fmt.Printf(" Kick off processing file %d name %s\n", i, full_file_name)
        go process_file(full_file_name, i, &wg)
    }

    /* Wait until all GoRoutines are done */
    wg.Wait()

    /* Now iterate over results for each file
     * accumulate the map from results[i] into the master_results map
     */

    for i, _ := range files_info {
        
        for word, word_count := range results[i] {
            master_results[word] += word_count 
        }
    }

    for word, word_count := range master_results {
        fmt.Printf("%v %v \n", word, word_count)
        fmt.Fprintf(fo, "%v %v \n", word, word_count)
    }

}

