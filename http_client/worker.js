var time = performance.now();

const sleep = (ms) => {
    return new Promise((resolve) => {
      setTimeout(resolve, ms);
    });
}

const jobsCount = 4;
const workerCount = 2;

var jobs = []
var results = []

for (let i = 0; i < jobsCount; i++)  {
    jobs.push(i+1)
    console.log("jobs", i+1)

}

const worker = (id, jobs, results) => {
    jobs.map( async (_, index) => {
        await sleep(1000)
		console.log("worker #%d finished\n", id)
		results.push(index * index)
    })
}

for (let i = 0; i < workerCount; i++) {
    worker(i+1, jobs, results)
}

for (let i = 0; i < jobsCount; i++)  {
    console.log("result #%d : value %s \n", i+1, results[i])    
}

time = performance.now() - time;
console.log('Время выполнения = ', time);


