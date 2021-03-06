let stock_log_data;

function update_stock_log_table(data) {
    const tbody = document.getElementById("stockLogsTableBody");
    tbody.innerHTML = "";
    for (var i=0; i<data.length; i++) {
        var p = "";
        p += "<tr>";
        p += "<td>" + data[i].log_id + "</td>";
        p += "<td>" + data[i].stock_id + "</td>";
        p += "<td>" + data[i].prod_id + "</td>";
        p += "<td>" + data[i].qty + "</td>";
        p += "<td>" + data[i].date_processed + "</td>";
        p += "<td>" + data[i].expiry_date + "</td>";
        p += "<td>" + data[i].action + "</td>";
        p += "<td>" + data[i].status + "</td>";
        p += "</tr>";
        tbody.insertAdjacentHTML("beforeend", p);
    }
}

async function get_stock_logs() {
    let res = await fetch("/stock_log", {
        "method": "GET",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        }
    })
    if (res.status === 401) {
        window.location.href = "/"
    }
    let data = await res.json()
    return data["data"]
}

const stockLogView = document.getElementById("stock-log-view")
stockLogView.addEventListener("change",(e) => {
    getStockLogBy(e.target.value)
})

async function getStockLogBy(type) {
    var stmt
    if (type === "-1") {
        stmt = "/all"
    } else {
        stmt = "/".concat(type);
    }
    let res = await fetch("/stock_log".concat(stmt), {
        "method": "GET",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        }
    })
    if (res.status === 401) {
        window.location.href = "/"
    }
    let data = await res.json()
    stock_log_data = data["data"]
    update_stock_log_table(stock_log_data)
}

async function init_stock_logs() {
    stock_log_data = await get_stock_logs();
    stock_log_data.sort((a, b) => {
        return b.log_id-a.log_id;
    })
    update_stock_log_table(stock_log_data)
}

init_stock_logs()
