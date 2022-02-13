let product_data;

function update_products_table(data) {
    const tbody = document.getElementById("productsTableBody");
    tbody.innerHTML = "";
    for (var i=0; i<data.length; i++) {
        var p = "";
        p += "<tr>";
        p += "<td>" + data[i].prod_id + "</td>";
        p += "<td>" + data[i].prod_name + "</td>";
        p += "<td>" + data[i].cat_name + "</td>";
        p += "<td>" + data[i].rate + "</td>";
        p += "<td>" + data[i].total_qty + "</td>";
        p += "<td>" + data[i].max_capacity + "</td>";
        p += "<td>" + data[i].s_name + "</td>";
        p += "</tr>";
        tbody.insertAdjacentHTML("afterend", p);
    }
}

async function get_products() {
    let res = await fetch("http://localhost/product", {
        "method": "GET",
        "headers": {}
    })
    let data = await res.json()
    return data["data"]
}

async function init() {
    product_data = await get_products();
    update_products_table(product_data)
}

init()
