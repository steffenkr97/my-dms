
// // Event-Handler für Formularübermittlung
// document.getElementById('upload-form').addEventListener('htmx:configRequest', function (event) {
//   // Hier kannst du vor dem Senden der Anfrage zusätzliche Anpassungen vornehmen, falls erforderlich.
//   // Du kannst event.detail.xhr verwenden, um auf das XMLHttpRequest-Objekt zuzugreifen.
//
//   // Zum Beispiel, wenn du zusätzliche Header hinzufügen möchtest:
//   event.detail.xhr.setRequestHeader('Authorization', 'Bearer YourAccessToken');
// });

document.getElementById('upload-form').addEventListener('htmx:response', function (event) {
  // Diese Funktion wird aufgerufen, nachdem die Antwort vom Server empfangen wurde.
  // Du kannst die Antwort verarbeiten oder weitere Aktionen ausführen.

  if (event.detail.verb === 'post') {
    const responseElement = document.getElementById('response');
    if (event.detail.status === 200) {
      // Erfolgreiche Übermittlung
      responseElement.innerText = 'Dokument erfolgreich hochgeladen.'; 

      window.alert("erfolgreich")
      
      // Redirect zur Startseite ("/")
      window.location.href = '/';

    } else {
      // Fehler bei der Übermittlung
      responseElement.innerText = 'Fehler beim Hochladen des Dokuments: ' + event.detail.statusText;
    }
  }
});
