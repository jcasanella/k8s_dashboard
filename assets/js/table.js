async function fetchAlbumsJSON() {
    const response = await fetch("http://localhost:8085/v1/albums");

    if (!response.ok) {
        const message = 'An error has occured: ${response.status}';
        throw new Error(message);
    }

    const albums = await response.json();
    return albums;
}

function getAlbums() {
    fetchAlbumsJSON().then(data => {
        if (Object.keys(data).length > 0) {

            var table_json = document.getElementById("table_json");
            var table = document.createElement("table");
            table.classList.add("table", "table-striped", "table-hover");

            createHeader(table);

            var tBody = document.createElement("tbody");
            for (i = 0; i < data.length; i++) {
                createRow(data[i], tBody);
            }
            table.appendChild(tBody);
            table_json.appendChild(table);
        }
        console.log(data);
    }).catch(error => {
        alert(error.message);
    });
}

function createHeader(table) {
    var tHead = document.createElement("thead");
    var row = document.createElement("tr");

    createHeaderCell(row, "Id");
    createHeaderCell(row, "Name");
    createHeaderCell(row, "Title");
    createHeaderCell(row, "Price");

    tHead.appendChild(row);
    table.appendChild(tHead);
}

function createHeaderCell(row, text) {
    var th = document.createElement("th");
    th.setAttribute("scope", "col");
    var cell = document.createTextNode(text);
    th.appendChild(cell);
    row.appendChild(th);
}

function createRow(text, table) {
    var row = document.createElement("tr");

    createCell(text.id, row);
    createCell(text.artist, row);
    createCell(text.title, row);
    createCell(text.price, row);

    table.appendChild(row);
}

function createCell(text, row) {
    var cell = document.createElement("td");
    var textCell = document.createTextNode(text);
    cell.appendChild(textCell);
    row.appendChild(cell);
}