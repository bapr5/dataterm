const term = document.getElementById('terminal');
const cmdline = document.getElementById('textline');

document.addEventListener("keydown", function(event)
{
     if (event.key == "Enter"){
        handleCommand(textline.value);
        textline.value="";
     }
});

function handleCommand(command) {
  switch(command) {
    case 'help':
      typeCommand(command,'List of available commands:');
      typeCommand(command,'help - get this list');
      typeCommand(command,'about - info about this site');
      typeCommand(command,'contacts - get contacts info');
      break;
    case 'about':
      typeCommand(command,'Это сайт, созданный для демонстрации навигации в терминале.');
      break;
    case 'contact':
      typeCommand(command,'Свяжитесь с нами по адресу example@example.com');
      break;
    default:
      typeCommand(command,"dataterm: "+command+": No such command");
      break;
  }
}
function typeCommand(command,response){
  var p = document.createElement('p');
  p.appendChild(document.createTextNode(command));
  term.appendChild(p)
  var p = document.createElement('p');
  p.appendChild(document.createTextNode(response));
  term.appendChild(p)
}

