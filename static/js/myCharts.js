trendsInfo = $('#trendsInfo').data('info').slice(3);

var numDays = 31
var trendDataSets = [];
eachTrend = trendsInfo.match(/(\w)*:\{\[(\d|\s)*]}/g);
for (var i in eachTrend) {
    var dataMap = {};
    dataMap["fillColor"] = "rgba(220,220,220,0.5)";
    dataMap["strokeColor"] = "rgba(220,220,220,1)";
    dataMap["pointColor"] = "rgba(220,220,220,1)";
    dataMap["pointStrokeColor"] = "#fff";
    // name = "Scotland_Yard:"
    var name = eachTrend[i].match(/(\w)*:/g);
    // dataStr = "0 7 4 3 1 4 5 2 7 3 2 1 8 5 9 4 4 1 0 4 3 4 10 8 10 4 8 3 9 3 11"
    var dataStr = eachTrend[i].match(/(\d)+/g);
    
    var dataArray = new Array();
    if(dataStr["30"] > 0) {
        numDays = 30;
    }
    for(var j in dataStr) {
        dataArray[j] = parseInt(dataStr[j]);
    }
    dataMap['data'] = dataArray;
    trendDataSets.push(dataMap)
}

var daysIndex = new Array()
for (var i = 0; i <= numDays; i++) {
    daysIndex[i] = i.toString();
}

var lineChartData = {
    labels: daysIndex,
    datasets: trendDataSets
    // datasets: [{
    //     fillColor: "rgba(220,220,220,0.5)",
    //     strokeColor: "rgba(220,220,220,1)",
    //     pointColor: "rgba(220,220,220,1)",
    //     pointStrokeColor: "#fff",
    //     data: [65, 59, 90, 81, 56, 55, 40]
    // }, {
    //     fillColor: "rgba(151,187,205,0)",
    //     strokeColor: "rgba(151,187,205,1)",
    //     pointColor: "rgba(151,187,205,1)",
    //     pointStrokeColor: "#fff",
    //     data: [28, 48, 40, 19, 96, 27, 100]
    // }]

};

new Chart(document.getElementById("trendsLineChart").getContext("2d")).Line(lineChartData);


function showMyTrend() {
    document.write("fsfsd");
}