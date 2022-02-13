let product_data;
let supplier_list;
let category_list;

function openOperationWindow(index) {
    var openelement=document.getElementsByClassName("tableoperationwindow");
    openelement[0].style.display="inline-block";
    var x=document.getElementsByClassName("tableoperation");
    x[index].style.display="inline-block";
    get_suppliers_list();
    get_category_list();
}

function closeOperationWindow() {
    var closeelement=document.getElementsByClassName("tableoperationwindow");
    closeelement[0].style.display="none";
    closeelement[0].style.animation="fadeEffect .8s";
    var x=document.getElementsByClassName("tableoperation");
    for(i=0;i<x.length;i++){
        x[i].style.display="none";
    }
}

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
        tbody.insertAdjacentHTML("beforeend", p);
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

async function get_suppliers_list() {
    let res = await fetch("http://localhost/supplier", {
        "method": "GET",
        "headers": {}
    })
    let data = await res.json()
    supplier_list = data["data"]
    const select = document.getElementById("new-product-supplier");
    select.innerHTML = "";
    for (var i=0; i<supplier_list.length; i++) {
        var p= "";
        p += "<option value=\""+supplier_list[i].supplier_id+"\">"+supplier_list[i].s_name+"</option>";
        select.insertAdjacentHTML("beforeend", p);
    }
}

async function get_category_list() {
    let res = await fetch("http://localhost/product_category", {
        "method": "GET",
        "headers": {}
    })
    let data = await res.json()
    category_list = data["data"]
    const select = document.getElementById("new-product-category");
    select.innerHTML = "";
    for (var i=0; i<category_list.length; i++) {
        var p= "";
        p += "<option value=\""+category_list[i].cat_id+"\">"+category_list[i].cat_name+"</option>";
        select.insertAdjacentHTML("beforeend", p);
    }
}

async function init_products() {
    product_data = await get_products();
    product_data.sort((a, b) => {
        if ( a.prod_id < b.prod_id ){
            return -1;
        }
        if ( a.prod_id > b.prod_id ){
            return 1;
        }
        return 0;
    })
    update_products_table(product_data)
}

async function add_product(){
    productId=document.getElementById("new-product-id").value;
    productName=document.getElementById("new-product-name").value;
    productCategory=Number(document.getElementById("new-product-category").value);
    productRate=Number(document.getElementById("new-product-rate").value);
    productMaxCapacity=Number(document.getElementById("new-product-max-capacity").value);
    productSupplier=Number(document.getElementById("new-product-supplier").value);
    const res = await fetch("http://localhost/product", {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            prod_id: productId,
            prod_name: productName,
            supplier_id: productSupplier,
            cat_id: productCategory,
            rate: productRate,
            max_capacity: productMaxCapacity
        })
    })
    const data = res.json()
    console.log(data)
    init_products()
    closeOperationWindow()
}

async function new_category(){
    var categoryName=document.getElementById("new-category-name").value;
    var categoryLocation=document.getElementById("new-category-location").value;
    const res = await fetch("http://localhost/product_category", {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            cat_name: categoryName,
            warehouse_loc: categoryLocation
        })
    })
    const data = await res.json()
    console.log(data)
    get_category_list()
    closeOperationWindow()
}

async function add_stock(){
    newStockStockId=document.getElementById("newstock-stockid").value;
    newStockProductId=document.getElementById("newstock-prodid").value;
    newStockExpDate=document.getElementById("newstock-expdate").value;
    newStockQty=Number(document.getElementById("newstock-quantity").value);
    const res = await fetch("http://localhost/stock", {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            stock_id: newStockStockId,
            prod_id: newStockProductId,
            curr_qty: newStockQty,
            expiry_date: newStockExpDate
        })
    })
    const data = await res.json()
    console.log(data)
    init_products()
}

async function remove_stock(){
    removeStockProdId=document.getElementById("rmstock-prodid").value;
    removeStockQty=Number(document.getElementById("rmstock-quantity").value);
    for(var i=0; i<product_data.length; i++) {
        if (product_data[i].prod_id == removeStockProdId && product_data[i].total_qty < removeStockQty) {
            alert("Error - quantity too high")
            return
        }
    }
    const res = await fetch("http://localhost/stock/".concat(removeStockProdId, "/", removeStockQty), {
        "method": "DELETE",
        "headers": {}
    })
    const data = await res.json()
    console.log(data)
    init_products()
}

init_products()
get_suppliers_list()
get_category_list()
