document.addEventListener("DOMContentLoaded", function() {
    const endpoints = document.querySelectorAll(".endpoint");

    const indexData = document.getElementById("indexData");
    const fetchDataButton = document.getElementById("fetchDataButton");

    const getAllMembers = document.getElementById("getAllMembers");
    const fetchDataButton2 = document.getElementById("fetchDataButton2");

    const memberIDInput = document.getElementById("memberIDInput");
    const getMember = document.getElementById("getMember");
    const fetchDataButton3 = document.getElementById("fetchDataButton3");

    const memberDataInput = document.getElementById("memberDataInput");
    const postMemberResponse = document.getElementById("postMemberResponse");
    const postDataButton = document.getElementById("postDataButton");

    const updateMemberIDInput = document.getElementById("updateMemberIDInput");
    const updateMemberDataInput = document.getElementById("updateMemberDataInput");
    const updateMemberResponse = document.getElementById("updateMemberResponse");
    const updateDataButton = document.getElementById("updateDataButton");

    const deleteMemberForm = document.getElementById("deleteMemberForm");
    const deleteMemberIDInput = document.getElementById("deleteMemberIDInput");
    const deleteMemberResponse = document.getElementById("deleteMember");
    const deleteDataButton = document.getElementById("deleteDataButton");


    endpoints.forEach(endpoint => {
      endpoint.querySelector("h3").addEventListener("click", function() {
        this.nextElementSibling.classList.toggle("details");
      });
    });

    function fetchData() {
      fetch(`http://127.0.0.1:8080/`)
        .then(response => response.text())
        .then(data => {
          indexData.textContent = data;
        })
        .catch(error => {
          indexData.textContent = "Error fetching data: " + error;
        });
    }

    function fetchData2() {
      fetch(`http://127.0.0.1:8080/members`)
        .then(response => response.text())
        .then(data => {
          getAllMembers.textContent = data;
        })
        .catch(error => {
          getAllMembers.textContent = "Error fetching data: " + error;
        });
    }

    function fetchData3() {
      const memberID = memberIDInput.value;
      const url = `http://127.0.0.1:8080/members/${memberID}`;
      fetch(url)
        .then(response => response.text())
        .then(data => {
          getMember.textContent = data;
        })
        .catch(error => {
          getMember.textContent = "Error fetching data: " + error;
        });
    }

    function postData() {
        const memberData = memberDataInput.value;
        const url = "http://127.0.0.1:8080/members";

        fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: memberData
        })
        .then(response => response.text())
        .then(data => {
          postMemberResponse.textContent = data;
        })
        .catch(error => {
          postMemberResponse.textContent = "Error posting data: " + error;
        });
    }

    function updateData() {
        const memberID = updateMemberIDInput.value;
        const memberData = updateMemberDataInput.value;
        const url = `http://127.0.0.1:8080/members/${memberID}`;

        fetch(url, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json"
          },
          body: memberData
        })
          .then(response => response.text())
          .then(data => {
            updateMemberResponse.textContent = data;
          })
          .catch(error => {
            updateMemberResponse.textContent = "Error updating data: " + error;
          });
    }

    function deleteData() {
      const memberID = deleteMemberIDInput.value;
      const url = `http://127.0.0.1:8080/members/${memberID}`;

      fetch(url, {
        method: "DELETE",
      })
        .then(response => response.text())
        .then(data => {
          deleteMemberResponse.textContent = data;
        })
        .catch(error => {
          deleteMemberResponse.textContent = "Error deleting data: " + error;
        });
    }

    fetchDataButton.addEventListener("click", fetchData);
    fetchDataButton2.addEventListener("click", fetchData2);
    fetchDataButton3.addEventListener("click", fetchData3);
    postDataButton.addEventListener("click", postData);
    updateDataButton.addEventListener("click", updateData);
    deleteDataButton.addEventListener("click", deleteData);
});
