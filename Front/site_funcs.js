const apiUrl = "http://localhost:10000/Filter";

document.addEventListener("DOMContentLoaded", function() {
    const button_value = document.getElementById("filt_button").addEventListener("click", async function() {
        const date_init = document.getElementById("data_init").value;
        const date_end = document.getElementById("data_end").value;
        const materials = document.getElementById("materials").value;
        const color = document.getElementById("color").value;
        try {
            const params = new URLSearchParams();
            params.append('date_init', date_init)
            params.append('date_end',date_end);
            params.append('materials',materials)
            params.append('color',color)
            const queryString = params.toString();
            const url = `${apiUrl}?${queryString}`;
            const response = await fetch(url, {
                method: 'GET',
                mode: 'cors',
            });
            if (!response.ok) {
                throw new Error(`Response status: ${response.status}`);
            }
            const json = await response.json();
            console.log(json);
            const result = document.getElementById("std_dev");
            result.innerHTML = json;
        } catch (error) {
            console.error(error.message);
        }
    });
});
