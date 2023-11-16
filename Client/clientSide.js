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

  
  