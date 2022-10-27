async function custom(c) {
  const { value: result } = await Swal.fire({
    title: "Check Availability",
    html:
      '<label for="stdate" class="form-label required">Start Date</label>' +
      '<input type="date" id="stdate" class="swal2-input"><br />' +
      '<label for="endate" class="form-label required">End Date</label>' +
      '<input type="date" id="endate" class="swal2-input">',
    focusConfirm: false,
    backdrop: false,
    showCancelButton: true,
    preConfirm: () => {
      return [
        document.getElementById("stdate").value,
        document.getElementById("endate").value,
      ];
    },
  });

  if (result) {
    if (result.dismiss !== Swal.DismissReason.cancel) {
      if (result.value !== "") {
        document.getElementById("sdate").value = result[0];
        document.getElementById("edate").value = result[1];

        let form = document.getElementById("availabilityJSON");
        let formData = new FormData(form);
        formData.append("name", "Hossein");
        formData.append("password", "Hossein123");

        fetch("/availabilityJSON", {
          method: "post",
          body: formData,
        })
          .then((Response) => Response.json())
          .then((data) => {
            console.log(data);
            console.log(data.message);
          });
      } else {
        console.log(false);
      }
    } else {
      console.log(false);
    }
  }
}
