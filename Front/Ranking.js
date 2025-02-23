const serverUrl = "http://localhost:10000/";
async function return_color_avgs() {
  const material = document.getElementById("materials").value;
  try {
    const params = new URLSearchParams();
    params.append("materials", material);
    const queryString = params.toString();
    const url = `${serverUrl}ranking?${queryString}`;
    const response = await fetch(url, {
      method: "GET",
      mode: "cors",
    });
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }
    const json = await response.json();
    console.log(json);
    let configurado_list = document.getElementById("color_list");
    for (i = 0; i < json.length; ++i) {
      let li = document.createElement("li");
      li.innerText = json[i].configurado;
      configurado_list.appendChild(li);
    }
    let cp_list = document.getElementById("cp_list");
    for (i = 0; i < json.length; ++i) {
      let cp_li = document.createElement("cp_li");
      cp_li.innerText = json[i].value;
      cp_list.appendChild(cp_li);
    }
  } catch (error) {
    console.error(error.message);
  }
}
document.addEventListener("DOMContentLoaded", function () {
  const button_value = document.addEventListener("change", async function () {
    return_color_avgs();
  });
});
