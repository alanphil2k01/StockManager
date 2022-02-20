let supplier_data;

function update_supplier_table(data) {
    const tbody = document.getElementById("suppliersTableBody");
    tbody.innerHTML = "";
    for (var i=0; i<data.length; i++) {
        var p = "";
        p += "<tr>";
        p += "<td>" + data[i].supplier_id + "</td>";
        p += "<td>" + data[i].s_name + "</td>";
        p += "<td>" + data[i].address + "</td>";
        p += "<td>" + data[i].s_email + "</td>";
        p += "<td>" + data[i].manager + "</td>";
        p += "<td>" + data[i].phone_no + "</td>";
        p += "</tr>";
        tbody.insertAdjacentHTML("beforeend", p);
    }
}

async function get_suppliers() {
    let res = await fetch("/supplier", {
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

async function init_suppliers() {
    supplier_data = await get_suppliers();
    supplier_data.sort((a, b) => {
        return a.supplier_id-b.supplier_id;
    })
    update_supplier_table(supplier_data)
}

async function add_supplier() {
    sForm = document.getElementsByClassName("supplier-detail-form")[0]
    const res = await fetch("/supplier", {
        "method": "POST",
        "headers": {
            "Authorization": "Bearer " + window.localStorage.getItem('ssmc-jwt'),
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            s_name: sForm[0].value,
            s_email: sForm[1].value,
            manager: sForm[2].value,
            address: sForm[3].value,
            phone_no: sForm[4].value
        })
    })
    if (res.status === 401) {
        alert("Unauthorized")
    } else if (res.status === 400) {
        alert("Invalid input")
    } else if (res.status === 500) {
        alert("Couldn't add supplier. Supplier name already exists")
    } else {
        const data = await res.json();
        console.log(data)
        init_suppliers()
    }
}

init_suppliers()
