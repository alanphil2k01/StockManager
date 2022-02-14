let stock_data;
let temp_stock_data;

function update_stock_table(data) {
    const tbody = document.getElementById("stocklistTableBody");
    tbody.innerHTML = "";
    for (var i=0; i<data.length; i++) {
        var p = "";
        p += "<tr>";
        p += "<td>" + data[i].stock_id + "</td>";
        p += "<td>" + data[i].prod_id + "</td>";
        p += "<td>" + data[i].prod_name + "</td>";
        p += "<td>" + data[i].curr_qty + "</td>";
        p += "<td>" + data[i].expiry_date + "</td>";
        p += "</tr>";
        tbody.insertAdjacentHTML("beforeend", p);
    }
}

async function get_stocks() {
    let res = await fetch("http://localhost/stock", {
        "method": "GET",
        "headers": {}
    })
    let data = await res.json()
    return data["data"]
}

async function init_stocks() {
    stock_data = await get_stocks();
    stock_data.sort((a, b) => {
        return a.stock_id-b.stock_id;
    })
    update_stock_table(stock_data)
}

const searchInput = document.getElementById("stock-search");
searchInput.addEventListener("input", (e)=>{
    filterByName(e.target.value)
})

function filterByName(name) {
    if(!name) {
        update_stock_table(stock_data)
        return
    }
    temp_stock_data = stock_data.filter((p, _) => p.prod_name.toLowerCase().includes(name.toLowerCase()))
    update_stock_table(temp_stock_data)
}

init_stocks()
