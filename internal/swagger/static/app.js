document.addEventListener("DOMContentLoaded", function() {
    const endpoints = document.querySelectorAll(".endpoint");

    const indexButton = document.getElementById("indexButton");

    const getAllMembersButton = document.getElementById("getAllMembersButton");

    const getMemberInput = document.getElementById("getMemberInput");
    const getMemberButton = document.getElementById("getMemberButton");

    const addMemberInput = document.getElementById("addMemberInput");
    const addMemberButton = document.getElementById("addMemberButton");

    const updateMemberIDInput = document.getElementById("updateMemberIDInput");
    const updateMemberDataInput = document.getElementById("updateMemberDataInput");
    const updateMemberButton = document.getElementById("updateMemberButton");

    const deleteMemberInput = document.getElementById("deleteMemberInput");
    const deleteMemberButton = document.getElementById("deleteMemberButton");

    const getAllBooksButton = document.getElementById("getAllBooksButton");

    const getBookInput = document.getElementById("getBookInput");
    const getBookButton = document.getElementById("getBookButton");

    const addBookInput = document.getElementById("addBookInput");
    const addBookButton = document.getElementById("addBookButton");

    const updateBookIDInput = document.getElementById("updateBookIDInput");
    const updateBookDataInput = document.getElementById("updateBookDataInput");
    const updateBookButton = document.getElementById("updateBookButton");

    const deleteBookInput = document.getElementById("deleteBookInput");
    const deleteBookButton = document.getElementById("deleteBookButton");

    endpoints.forEach(endpoint => {
      endpoint.querySelector("h3").addEventListener("click", function() {
        this.nextElementSibling.classList.toggle("details");
      });
    });

    function fetchData(url, method, body = null) {
      return fetch(url, {
        method,
        headers: {
          "Content-Type": "application/json",
        },
        body,
      })
        .then((response) => Promise.all([response.text(), response.status, response.statusText]))
        .then(([data, status, statusText]) => ({ data, status, statusText }))
        .catch((error) => ({ error }));
    }

    function updateElement(dataElementId, statusElementId, data, status, statusText) {
      const dataElement = document.getElementById(dataElementId);
      const statusElement = document.getElementById(statusElementId);
      dataElement.textContent = `Data: ${data}`;
      statusElement.textContent = `Status: ${status} (${statusText})`;
      statusElement.setAttribute("data-status", status);
    }

    function indexFunc() {
      fetchData("http://127.0.0.1:8080/", "GET")
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const indexDataElement = document.getElementById("indexData");
            indexDataElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataIndexData", "statusIndexData", data, status, statusText);
          }
        });
    }

    // * Member
    function getAllMembersFunc() {
      fetchData("http://127.0.0.1:8080/members", "GET")
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const getAllMembersElement = document.getElementById("getAllMembers");
            getAllMembersElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataGetAllMembers", "statusGetAllMembers", data, status, statusText);
          }
        });
    }

    function getMemberFunc() {
      const memberID = getMemberInput.value;
      const url = `http://127.0.0.1:8080/members/${memberID}`;
      fetchData(url, "GET")
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const getMemberElement = document.getElementById("getMember");
            getMemberElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataGetMember", "statusGetMember", data, status, statusText);
          }
        });
    }

    function addMemberFunc() {
      const memberData = addMemberInput.value;
      const url = "http://127.0.0.1:8080/members";
      fetchData(url, "POST", memberData)
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const addMemberElement = document.getElementById("addMember");
            addMemberElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataAddMember", "statusAddMember", data, status, statusText);
          }
        });
    }

    function updateMemberFunc() {
        const memberID = updateMemberIDInput.value;
        const memberData = updateMemberDataInput.value;
        const url = `http://127.0.0.1:8080/members/${memberID}`;

        fetchData(url, "PUT", memberData)
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const updateMemberElement = document.getElementById("updateMember");
            updateMemberElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataUpdateMember", "statusUpdateMember", data, status, statusText);
          }
        });
    }

    function deleteMemberFunc() {
      const memberID = deleteMemberInput.value;
      const url = `http://127.0.0.1:8080/members/${memberID}`;

      fetchData(url, "DELETE")
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const deleteMemberElement = document.getElementById("deleteMember");
            deleteMemberElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataDeleteMember", "statusDeleteMember", data, status, statusText);
          }
        });
    }

    // * Book
    function getAllBooksFunc() {
      fetchData("http://127.0.0.1:8080/books", "GET")
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const getAllBooksElement = document.getElementById("getAllBooks");
            getAllBooksElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataGetAllBooks", "statusGetAllBooks", data, status, statusText);
          }
        });
    }

    function getBookFunc() {
      const bookID = getBookInput.value;
      const url = `http://127.0.0.1:8080/books/${bookID}`;

      fetchData(url, "GET")
      .then(({ data, status, statusText, error }) => {
        if (error) {

          const getBookElement = document.getElementById("getBook");
          getBookElement.textContent = "Error fetching data: " + error;
        } else {
          updateElement("dataGetBook", "statusGetBook", data, status, statusText);
        }
      });
    }

    function addBookFunc() {
      const bookData = addBookInput.value;
      const url = "http://127.0.0.1:8080/books";
      fetchData(url, "POST", bookData)
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const addBookElement = document.getElementById("addBook");
            addBookElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataAddBook", "statusAddBook", data, status, statusText);
          }
        });
    }

    function updateBookFunc() {
        const bookID = updateBookIDInput.value;
        const bookData = updateBookDataInput.value;
        const url = `http://127.0.0.1:8080/books/${bookID}`;

        fetchData(url, "PUT", bookData)
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const updateBookElement = document.getElementById("updateBook");
            updateBookElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataUpdateBook", "statusUpdateBook", data, status, statusText);
          }
        });
    }

    function deleteBookFunc() {
      const bookID = deleteBookInput.value;
      const url = `http://127.0.0.1:8080/books/${bookID}`;

      fetchData(url, "DELETE")
        .then(({ data, status, statusText, error }) => {
          if (error) {
            const deleteBookElement = document.getElementById("deleteBook");
            deleteBookElement.textContent = "Error fetching data: " + error;
          } else {
            updateElement("dataDeleteBook", "statusDeleteBook", data, status, statusText);
          }
        });
    }

    indexButton.addEventListener("click", indexFunc);
    getAllMembersButton.addEventListener("click", getAllMembersFunc);
    getMemberButton.addEventListener("click", getMemberFunc);
    addMemberButton.addEventListener("click", addMemberFunc);
    updateMemberButton.addEventListener("click", updateMemberFunc);
    deleteMemberButton.addEventListener("click", deleteMemberFunc);
    getAllBooksButton.addEventListener("click", getAllBooksFunc);
    getBookButton.addEventListener("click", getBookFunc);
    addBookButton.addEventListener("click", addBookFunc);
    updateBookButton.addEventListener("click", updateBookFunc);
    deleteBookButton.addEventListener("click", deleteBookFunc);
});
