function verifyAnswer(userAnsElement,correctAns,event,exerciseID,exerciseOuter){
    event.preventDefault();
    console.log("UserAns ID=",userAnsElement)
    console.log("CorrectAns=",correctAns)
    console.log("Exercise id=",exerciseID)
    userAns = Number(document.getElementById(userAnsElement).value)

    correctAns = Number(correctAns)
    console.log("user answer:",userAns,"type: ",typeof userAns)
    console.log("server answer:",correctAns,"type: ",typeof correctAns)



    if(userAns === correctAns){
        markInput(exerciseOuter,true)
        setStatus = "Completed"
        
    }
    else{
        markInput(exerciseOuter,false)
        setStatus = "Attempted"
    }
    fetch(`/exercises/${exerciseID}`, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "status": setStatus })
    })
    .then(response => {
      
        if (response.ok) {
            
            console.log("sent put request to server")
        } else {
          
            throw new Error('error on put request');
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
}


function markInput(inputId, isValid) {
    // Get the input element
    var input = document.getElementById(inputId);

    // Check if the input exists to avoid errors
    if (input) {
        // Set the border color based on the validity of the answer
        input.style.borderColor = isValid ? 'green' : 'red';
        // Optionally, you can add more styles like border-width or border-style
        input.style.borderWidth = '2px';
        input.style.borderStyle = 'solid';
    } else {
        console.warn('Input with id "' + inputId + '" not found.');
    }
}




