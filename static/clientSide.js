function goToLesson(lessonId) {
    // Make an AJAX request to the server to get details for the lesson with lessonId
    fetch(`/api/lessons/${lessonId}`)
      .then(response => response.json())
      .then(lessonDetails => {
        // Do something with the lesson details, like display them in a modal or navigate to a detail page
        console.log(lessonDetails);
      })
      .catch(error => {
        console.error('Error fetching lesson details:', error);
      });
}
// Login Page
function LoginAuth(form,event) {
  console.log("Now checking for authentication clientside")
  event.preventDefault();
  var username = form.elements["username"].value;
  var password = form.elements["password"].value;

  console.log("username inputted: ",username)
  console.log("password inputted: ",password)

  fetch('/login', {
      method: 'POST',
      headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({ "username": username, "password": password })
  })
  .then(response => {
      if (response.ok) {
          return response.json();
      } else {
          throw new Error('Login failed');
      }
  })
  .then(data => {
      console.log(data.message);
      window.location.href = '/lessons'; // Redirect to the new page
  })
  .catch(error => {
      console.error('Error:', error);
  });
}


// Register Page
function registerUser(event, form) {
  event.preventDefault();
  var username = form.elements["usernameRegister"].value;
  var password1 = form.elements["passwordRegister1"].value;
  var password2 = form.elements["passwordRegister2"].value;

  errorString = validateFormInformationForRegistration(username, password1, password2)

  if (errorString != "") {
      console.error('Error: ', errorString);
      showError(errorString);
      return;
  }

  fetch('/register', {
      method: 'POST',
      headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      
      body: JSON.stringify({ "username": username, "password": password1 })
  })
  .then(response => {
    
      if (response.ok) {
          
          return response.json();
      } else {
        
          throw new Error('Registration Failed');
      }
  })
  .then(data => {
    
    let jwtToken = data.token; 


    localStorage.setItem('jwtToken', jwtToken);

    window.location.href = '/lessons';
    
})
  .catch(error => {
      console.error('Error:', error);
  });
}


function validateFormInformationForRegistration(username,password1,password2) {
  astring = ""
  if(password1 != password2){
    astring = "Passwords must match.";

  }

  if (password1.length < 7) {
      astring = "Password must be at least 7 characters long and contain at least 1 special character and 1 number.";
      
  }
  if(username.length < 5){
    astring = "Username must be at least 5 characters long";
    
  }

  return astring;
}

function markInputError(inputElement) {
  inputElement.classList.add('input-error');
}

function clearInputError(inputElement) {
  inputElement.classList.remove('input-error');
}




  
  