var options = {
        //Boolean - If we show the scale above the chart data    
        scaleOverlay : false,
        
        //Boolean - If we want to override with a hard coded scale
        scaleOverride : false,
        
        //** Required if scaleOverride is true **
        //Number - The number of steps in a hard coded scale
        scaleSteps : null,
        //Number - The value jump in the hard coded scale
        scaleStepWidth : null,
        //Number - The scale starting value
        scaleStartValue : null,

        //String - Colour of the scale line    
        scaleLineColor : "rgba(0,0,0,1)",
        
        //Number - Pixel width of the scale line
        scaleLineWidth : 1,

        //Boolean - Whether to show labels on the scale    
        scaleShowLabels : true,
        
        //Interpolated JS string - can access value
        scaleLabel : "<%=value%>",
        
        //String - Scale label font declaration for the scale label
        scaleFontFamily : "'Arial'",
        
        //Number - Scale label font size in pixels    
        scaleFontSize : 12,
        
        //String - Scale label font weight style
        scaleFontStyle : "normal",
        
        //String - Scale label font colour    
        scaleFontColor : "#666",    
        
        ///Boolean - Whether grid lines are shown across the chart
        scaleShowGridLines : true,
        
        //String - Colour of the grid lines
        scaleGridLineColor : "rgba(0,0,0,.1)",
        
        //Number - Width of the grid lines
        scaleGridLineWidth : 1,    
        
        //Boolean - Whether the line is curved between points
        bezierCurve : true,
        
        //Boolean - Whether to show a dot for each point
        pointDot : true,
        
        //Number - Radius of each point dot in pixels
        pointDotRadius : 1,
        
        //Number - Pixel width of point dot stroke
        pointDotStrokeWidth : 0.2,
        
        //Boolean - Whether to show a stroke for datasets
        datasetStroke : true,
        
        //Number - Pixel width of dataset stroke
        datasetStrokeWidth : 3,
        
        //Boolean - Whether to fill the dataset with a colour
        datasetFill : true,
        
        //Boolean - Whether to animate the chart
        animation : true,

        //Number - Number of animation steps
        animationSteps : 60,
        
        //String - Animation easing effect
        animationEasing : "easeOutQuart",

        //Function - Fires when the animation is complete
        onAnimationComplete : null
    }

trendsInfo = $('#topicsTrendsInfo').data('info');


var numDays = 31;
var trendDataSets = [];
var IdSet = [];

eachTrend = trendsInfo.match(/\[(\d|\s|e|\.|\-|\+)*\]/g);

for (var i in eachTrend) {
    var dataMap = {};
    dataMap["fillColor"] = "rgba(220,220,220,0.4)";
    dataMap["strokeColor"] = "rgba(220,220,220,0.8)";
    dataMap["pointColor"] = "rgba(220,220,220,0.8)";

    IdSet.push(i)
    // dataStr = "0 7 4 3 1 4 5 2 7 3 2 1 8 5 9 4 4 1 0 4 3 4 10 8 10 4 8 3 9 3 11"
    var dataStr = eachTrend[i].match(/(\d|\.|e|\-)+/g);

    var dataArray = new Array();
    if (dataStr["30"] > 0) {
        numDays = 30;
    }
    for (var j in dataStr) {
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
};

var chart = new Chart(document.getElementById("topicsTrendsLineChart").getContext("2d")).Line(lineChartData, options);

function showMyTopicTrend(topicDataIndex) {
    dataId = -1;
    for (var i = 0; i < IdSet.length; i++) {
        if (topicDataIndex == IdSet[i]) {
            dataId = i;
            break;
        }
    }
    tmpDatasets = lineChartData["datasets"]
    for (var i = 0; i < tmpDatasets.length; i++) {
        tmpDatasets[i]["fillColor"] = "rgba(220,220,220,0.1)";
        tmpDatasets[i]["strokeColor"] = "rgba(220,220,220,0.1)";
        tmpDatasets[i]["pointColor"] = "rgba(220,220,220,0)";
    }

    lineChartData["datasets"][dataId]["fillColor"] = "rgba(0,0,0,1)";
    lineChartData["datasets"][dataId]["strokeColor"] = "rgba(0,0,0,1)";
    lineChartData["datasets"][dataId]["pointColor"] = "rgba(0,0,0,1)";

    chart = new Chart(document.getElementById("topicsTrendsLineChart").getContext("2d")).Line(lineChartData, options);
};


function showOurTrends(rowIndex, colIndex) {
    rowId = -1;
    colId = -1;
    for (var i = 0; i < IdSet.length; i++) {
        if (rowIndex == IdSet[i]) {
            rowId = i;
        }
        if (colIndex == IdSet[i]) {
            colId = i;
        }
    }
    tmpDatasets = lineChartData["datasets"]
    for (var i = 0; i < tmpDatasets.length; i++) {
        tmpDatasets[i]["fillColor"] = "rgba(220,220,220,0.1)";
        tmpDatasets[i]["strokeColor"] = "rgba(220,220,220,0.1)";
        tmpDatasets[i]["pointColor"] = "rgba(220,220,220,0)";
    }

    lineChartData["datasets"][rowId]["fillColor"] = "rgba(88,220,220,0.5)";
    lineChartData["datasets"][rowId]["strokeColor"] = "rgba(88,220,220,1)";
    lineChartData["datasets"][rowId]["pointColor"] = "rgba(88,220,220,1)";

    lineChartData["datasets"][colId]["fillColor"] = "rgba(151,187,205,0.5)";
    lineChartData["datasets"][colId]["strokeColor"] = "rgba(151,187,205,1)";
    lineChartData["datasets"][colId]["pointColor"] = "rgba(151,187,205,1)";

    chart = new Chart(document.getElementById("topicsTrendsLineChart").getContext("2d")).Line(lineChartData, options);
};
