let product_data;
let supplier_list;
let prod_cat_list;
let temp_product_data;

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
    let res = await fetch("/product", {
        "method": "GET",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        }
    })
    if (res.status === 401) {
        window.location.href = "/login.html";
        return
    }
    let data = await res.json()
    return data["data"]
}

async function get_suppliers_list() {
    let res = await fetch("/supplier", {
        "method": "GET",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        }
    })
    let data = await res.json()
    supplier_list = data["data"]
    const selectNew = document.getElementById("new-product-supplier");
    const selectEdit = document.getElementById("edit-product-supplier");
    selectNew.innerHTML = "";
    selectEdit.innerHTML = "";
    for (var i=0; i<supplier_list.length; i++) {
        var p= "";
        p += "<option value=\""+supplier_list[i].supplier_id+"\">"+supplier_list[i].s_name+"</option>";
        selectNew.insertAdjacentHTML("beforeend", p);
        selectEdit.insertAdjacentHTML("beforeend", p);
    }
}

function get_prod_id_list() {
    const selectEdit = document.getElementById("edit-product-id");
    selectEdit.innerHTML = "";
    for (var i=0; i<product_data.length; i++) {
        var p= "";
        p += "<option value=\""+product_data[i].prod_id+"\">"+product_data[i].prod_id+"</option>";
        selectEdit.insertAdjacentHTML("beforeend", p);
    }
    update_edit_form(0)
}

function update_edit_form(index) {
    document.getElementById('edit-product-id').value = product_data[index].prod_id
    document.getElementById('edit-product-category').value = product_data[index].cat_id
    document.getElementById('edit-product-supplier').value = product_data[index].supplier_id
    document.getElementById('edit-product-name').value = product_data[index].prod_name
    document.getElementById('edit-product-rate').value = product_data[index].rate
    document.getElementById('edit-product-max-capacity').value = product_data[index].max_capacity
    document.getElementById("edit-product-qty").innerHTML = product_data[index].total_qty
}

const editId = document.getElementById("edit-product-id")
editId.addEventListener("change", (e) => {
    for(let i=0; i<product_data.length; i++) {
        if(product_data[i].prod_id === e.target.value) {
            update_edit_form(i)
        }
    }
})


async function get_category_list() {
    let res = await fetch("/product_category", {
        "method": "GET",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        }
    })
    let data = await res.json()
    prod_cat_list = data["data"]
    const selectNew = document.getElementById("new-product-category");
    const selectEdit = document.getElementById("edit-product-category");
    selectNew.innerHTML = "";
    selectEdit.innerHTML = "";
    for (var i=0; i<prod_cat_list.length; i++) {
        var p= "";
        p += "<option value=\""+prod_cat_list[i].cat_id+"\">"+prod_cat_list[i].cat_name+"</option>";
        selectNew.insertAdjacentHTML("beforeend", p);
        selectEdit.insertAdjacentHTML("beforeend", p);
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
    temp_product_data = product_data
    update_products_table(product_data)
    get_prod_id_list()
}

const prodViewMenu = document.getElementById("prod-view");
prodViewMenu.addEventListener("change", (e)=>{
    filterByQty(e.target.value)
})

function filterByQty(type) {
    searchInput.value = ""
    if(type === "all") {
        temp_product_data = product_data
    } else if (type === "out_of_stock") {
        temp_product_data = product_data.filter((p, _) => p.total_qty === 0)
    } else if (type === "running_out") {
        temp_product_data = product_data.filter((p, _) => p.total_qty != 0 && ((p.total_qty/p.max_capacity)*100) < 5)
    }
    update_products_table(temp_product_data)
}

const searchInput = document.getElementById("prod-search");
searchInput.addEventListener("input", (e)=>{
    filterByName(e.target.value)
})

function filterByName(name) {
    if(!name) {
        update_products_table(product_data)
        return
    }
    temp_data = temp_product_data.filter((p, _) => p.prod_name.toLowerCase().includes(name.toLowerCase()))
    update_products_table(temp_data)
}

async function add_product(){
    productId=document.getElementById("new-product-id").value;
    productName=document.getElementById("new-product-name").value;
    productCategory=Number(document.getElementById("new-product-category").value);
    productRate=Number(document.getElementById("new-product-rate").value);
    productMaxCapacity=Number(document.getElementById("new-product-max-capacity").value);
    productSupplier=Number(document.getElementById("new-product-supplier").value);
    const res = await fetch("/product", {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
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
    if (res.status === 401) {
        alert('Unauthorized')
    } else if (res.status === 400) {
        alert("Invalid input")
    } else if (res.status === 500) {
        alert("Couldn't add product. Check if product id already exists.")
    } else if (res.status === 200) {
        const data = await res.json()
        console.log(data)
        init_products()
        closeOperationWindow()
    }
}

async function edit_product(){
    productId=document.getElementById("edit-product-id").value;
    productName=document.getElementById("edit-product-name").value;
    productCategory=Number(document.getElementById("edit-product-category").value);
    productRate=Number(document.getElementById("edit-product-rate").value);
    productMaxCapacity=Number(document.getElementById("edit-product-max-capacity").value);
    productSupplier=Number(document.getElementById("edit-product-supplier").value);
    const res = await fetch("/product/".concat(productId), {
        "method": "PUT",
        "headers": {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        },
        "body": JSON.stringify({
            prod_name: productName,
            supplier_id: productSupplier,
            cat_id: productCategory,
            rate: productRate,
            max_capacity: productMaxCapacity
        })
    })
    if (res.status === 401) {
        alert('Unauthorized')
    } else if (res.status === 500) {
        alert("Couldn't update")
        return
    } else if (res.status === 201){
        const data = await res.json()
        console.log(data)
        init_products()
        closeOperationWindow()
    }
}

async function delete_product(){
    productId=document.getElementById("edit-product-id").value;
    productTotalQty=Number(document.getElementById("edit-product-qty").value);
    if(productTotalQty > 0) {
        alert("Cannot remove product. Remove all stock for this product")
        return
    }
    const res = await fetch("/product/".concat(productId), {
        "method": "DELETE",
        "headers": {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        },
    })
    if (res.status === 401) {
        alert('Unauthorized')
    } else if (res.status === 500) {
        alert("Couldn't delete")
        return
    } else if (res.status === 200){
        const data = await res.json()
        console.log(data)
        init_products()
        closeOperationWindow()
    }
}

async function new_category(){
    var categoryName=document.getElementById("new-category-name").value;
    var categoryLocation=document.getElementById("new-category-location").value;
    const res = await fetch("/product_category", {
        "method": "POST",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            cat_name: categoryName,
            warehouse_loc: categoryLocation
        })
    })
    if (res.status === 401) {
        alert('Unauthorized')
    } else if (res.status === 400) {
        alert("Invalid input")
    } else if (res.status === 500) {
        alert("Couldn't create category. Already exists")
    } else if (res.status === 200) {
        const data = await res.json()
        console.log(data)
        get_category_list()
        closeOperationWindow()
    }
}

async function add_stock(){
    product_data = await get_products()
    newStockStockId=document.getElementById("newstock-stockid").value;
    newStockProductId=document.getElementById("newstock-prodid").value;
    newStockExpDate=document.getElementById("newstock-expdate").value;
    newStockQty=Number(document.getElementById("newstock-quantity").value);
    const res = await fetch("/stock", {
        "method": "POST",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            stock_id: newStockStockId,
            prod_id: newStockProductId,
            curr_qty: newStockQty,
            expiry_date: newStockExpDate
        })
    })
    if (res.status === 401) {
        alert('Unauthorized')
    }
    if (res.status === 400) {
        alert("Invalid input")
        return
    }
    const data = await res.json()
    console.log(data)
    init_products()
}

async function remove_stock(){
    product_data = await get_products()
    removeStockProdId=document.getElementById("rmstock-prodid").value;
    removeStockQty=Number(document.getElementById("rmstock-quantity").value);
    for(var i=0; i<product_data.length; i++) {
        if (product_data[i].prod_id == removeStockProdId && product_data[i].total_qty < removeStockQty) {
            alert("Error - quantity too high")
            return
        }
    }
    const res = await fetch("/stock/".concat(removeStockProdId, "/", removeStockQty), {
        "method": "DELETE",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
        }
    })
    if (res.status === 401) {
        alert('Unauthorized')
    }
    const data = await res.json()
    console.log(data)
    init_products()
}

init_products()
get_suppliers_list()
get_category_list()
