const serverUrl = "http://localhost:10000/";
async function insert() {
  const linhas = document.getElementById("linhas").value;
  const massa = document.getElementById("massa").value;
  const primer = document.getElementById("primer").value;
  const verniz = document.getElementById("verniz").value;
  const esmalte = document.getElementById("esmalte").value;
  const tingidor = document.getElementById("tingidor").value;
  const color = document.getElementById("color_list").value;
  const params = new URLSearchParams();
  params.append("linhas", linhas);
  params.append("massa", massa);
  params.append("primer", primer);
  params.append("verniz", verniz);
  params.append("esmalte", esmalte);
  params.append("tingidor", tingidor);
  params.append("color_list", color);
  const queryString = params.toString();
  const url = `${serverUrl}Insert?${queryString}`;
  const response = await fetch(url, {
    method: "GET",
    mode: "cors",
  });
  if (!response.ok) {
    throw new Error(`Response status: ${response.status}`);
  }
}
async function return_colors() {
  const response = await fetch(serverUrl + "dropdown", {
    mode: "cors",
  });
  if (!response.ok) {
    throw new Error(`Response status: ${response.status}`);
  }
  const json = await response.json();
  console.log(json);
  for (var i = 0; i < json.length; i++) {
    var opt = json[i];
    var el = document.createElement("option");
    el.textContent = opt;
    el.value = opt;
    document.getElementById("color_list").appendChild(el);
  }
}
document.addEventListener("DOMContentLoaded", function () {
  const button_value = document
    .getElementById("submit")
    .addEventListener("click", async function () {
      insert();
    });
  return_colors();
});