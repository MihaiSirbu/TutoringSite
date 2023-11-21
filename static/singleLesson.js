function verifyAnswer(userAns,correctAns,event,exerciseID){
    event.preventDefault();
    userAns = strconv.Atoi(userAns.value)

    correctAns = strconv.Atoi(correctAns)

    if(userAns === correctAns){
        // send POST fetch request to change status to completed
        // mark green
    }
    else{
        markInputError(document.getElementById(userAns))   
    }
}



function markInputError(inputElement) {
    inputElement.classList.add('input-error');
  }
  
  function clearInputError(inputElement) {
    inputElement.classList.remove('input-error');
  }
