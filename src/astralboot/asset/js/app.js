source = new EventSource('/events');
source.addEventListener('info',function(e){
    console.log(e);
})

source.onmessage = function(e) {
    console.log(e);
};
