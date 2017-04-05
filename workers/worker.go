package worker


func worker(name string, jobs <-chan map[string]interface{}, results chan<- map[string]interface{}) {
    for j := range jobs {
        j["task"]
        // run task
        results <- j
    }
}
