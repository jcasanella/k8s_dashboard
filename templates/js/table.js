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

            for (i = 0; i < data.length; i++) {
                createRow(data[i], table);
            }

            table_json.appendChild(table);
        }
        console.log(data);
    }).catch(error => {
        alert(error.message);
    });
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