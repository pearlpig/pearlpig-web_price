window.onload = function () {
    let oSearch = document.getElementById("search")
    oSearch.onclick = function () {
        let oWant = document.getElementById("want")
        let wanted = {"item": encodeURI(oWant.value) ,"server":["ruten"], "need_sort": true}
        let req =  JSON.stringify(wanted);
        $.ajax({
            type: 'GET', 
            url: '/result/',
            dataType: 'json',
            data: {"request":req},
            beforeSend: function () {
                // clear the previous result
                let oResult = document.getElementById("result")
                while (oResult.firstChild) {
                    oResult.removeChild(oResult.firstChild);
                }
            },
            success: function (data) {
                good = JSON.parse(JSON.stringify(data));
                // show the all result
                let cnt=0;
                data.forEach(item => {
                    add(item.info)
                    cnt++;
                });
            },
            error: function (e) {
                console.log(e);
            }
        });
    }
}

function add(data) {
    let oResult = document.getElementById("result")

    // url 
    let oA = document.createElement("a");
    let msg = data.url
    oA.setAttribute("href", msg)
    oA.setAttribute("target", "_blank")
    oA.appendChild(document.createTextNode(msg))
    oResult.appendChild(oA);

    // price
    let oP = document.createElement("p");
    msg = "價格" + data.price;
    oP.appendChild(document.createTextNode(msg))
    oResult.appendChild(oP);
}
