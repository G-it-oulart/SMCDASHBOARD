const serverUrl = "http://localhost:10000/";
async function build_graph() {
  const date_init = document.getElementById("data_init").value;
  const date_end = document.getElementById("data_end").value;
  const materials = document.getElementById("materials").value;
  const color = document.getElementById("color_list").value;
  try {
    const params = new URLSearchParams();
    params.append("date_init", date_init);
    params.append("date_end", date_end);
    params.append("materials", materials);
    params.append("color", color);
    const queryString = params.toString();
    const url = `${serverUrl}Filter?${queryString}`;
    const response = await fetch(url, {
      method: "GET",
      mode: "cors",
    });
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }
    const json = await response.json();
    console.log(json);
    if (window.myChart instanceof Chart) {
      window.myChart.destroy();
    }
    const ctx = document.getElementById("myChart");
    window.myChart = new Chart(ctx, {
      type: "line",
      data: {
        labels: json.Lab,
        datasets: [
          {
            label: "Actual Value",
            data: json.Val,
            borderWidth: 1,
          },
          {
            label: "UpperControl",
            data: json.Ucl,
            borderWidth: 1,
          },
          {
            label: "LowerControl",
            data: json.Lcl,
            borderWidth: 1,
          },
          {
            label: "Average",
            data: json.Avg,
            borderWidth: 1,
          },
          {
            label: "Standard",
            data: json.Std,
            borderWidth: 1,
          },
        ],
      },
      options: {
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      },
    });
  } catch (error) {
    console.error(error.message);
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
    .getElementById("filt_button")
    .addEventListener("click", async function () {
      build_graph();
    });
  return_colors();
  $("input[type=date]").keyup(function () {
    var datevalue = $(this).val();
    var dateSplit = datevalue.split("-"); // yyyy-mm-dd
    var dateYear = dateSplit[0]; // yyyy
    var dateMonth = dateSplit[1]; // mm
    var dateDay = dateSplit[2]; // dd
    if (dateYear.length > 4) {
      dateYear = dateYear.substring(0, 4);
      $(this).val(dateYear + "-" + dateMonth + "-" + dateDay);
    }
  });
});
