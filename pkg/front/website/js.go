// GENERATED BY textFileToGoConst
// GitHub:     github.com/logrusorgru/textFileToGoConst
// input file: html\front\website.js
// generated:  Thu Aug  4 15:52:25 CEST 2022

package website

const JS = "getWallets();\r\ngetWebsiteDeployerSC();\r\n\r\nlet wallets = [];\r\nlet deployers = [];\r\n\r\n// open file upload\r\nfunction openDialog(count) {\r\n\tconsole.log(count);\r\n\tdocument.getElementById('fileid0').value = null;\r\n\tdocument.getElementById('fileid0').click();\r\n}\r\n\r\n// Handle event on file selecting\r\nfunction handleFileSelect(evt) {\r\n\tlet files = evt.target.files; // get files\r\n\tlet f = files[0];\r\n\tconst reader = new FileReader();\r\n\treader.onload = (event) => importWallet(JSON.parse(event.target.result)); // desired file content\r\n\treader.onerror = (error) => reject(error);\r\n\treader.readAsText(f);\r\n}\r\n\r\nfunction errorAlert(error) {\r\n\tdocument.getElementsByClassName('alert-danger')[0].style.display = 'block';\r\n\r\n\tdocument.getElementsByClassName('alert-danger')[0].innerHTML = error;\r\n\r\n\tsetTimeout(function () {\r\n\t\tdocument.getElementsByClassName('alert-danger')[0].style.display = 'none';\r\n\t}, 5000);\r\n}\r\n\r\nfunction successMessage(message) {\r\n\tdocument.getElementsByClassName('alert-primary')[0].style.display = 'block';\r\n\r\n\tdocument.getElementsByClassName('alert-primary')[0].innerHTML = message;\r\n\r\n\tsetTimeout(function () {\r\n\t\tdocument.getElementsByClassName('alert-primary')[0].style.display = 'none';\r\n\t}, 5000);\r\n}\r\n// Import a wallet through PUT query\r\nasync function feedSelectOption(w) {\r\n\tconst formSelect = document.getElementsByClassName('form-select');\r\n\tfor (let i = 0; i < w.length; i++) {\r\n\t\tconst opt = document.createElement('option');\r\n\t\topt.value = i;\r\n\t\topt.text = w[i].nickname;\r\n\t\tformSelect[0].append(opt);\r\n\t}\r\n}\r\n\r\n// Create a wallet through POST query\r\nasync function getWallets() {\r\n\taxios\r\n\t\t.get('/mgmt/wallet')\r\n\t\t.then((resp) => {\r\n\t\t\tif (resp) {\r\n\t\t\t\tconst data = resp.data;\r\n\t\t\t\twallets = data;\r\n\t\t\t\tfeedSelectOption(wallets);\r\n\t\t\t}\r\n\t\t})\r\n\t\t.catch((e) => {\r\n\t\t\terrorAlert(e);\r\n\t\t});\r\n}\r\n\r\nasync function getWebsiteDeployerSC() {\r\n\taxios\r\n\t\t.get('/uploadWeb')\r\n\t\t.then((websites) => {\r\n\t\t\tlet count = 0;\r\n\t\t\tfor (const website of websites.data) {\r\n\t\t\t\ttableInsert(website, count);\r\n\t\t\t\tcount++;\r\n\t\t\t}\r\n\t\t\tdeployers = websites.data;\r\n\t\t})\r\n\t\t.catch((e) => {\r\n\t\t\terrorAlert(e.response.data.code);\r\n\t\t});\r\n}\r\n\r\nasync function deployWebsiteDeployerSC() {\r\n\tconst dnsNameInputValue = document.getElementById('websiteName').value;\r\n\r\n\tif (dnsNameInputValue == '') {\r\n\t\tconsole.log(dnsNameInputValue == '');\r\n\t\terrorAlert('Input a DNS name');\r\n\t} else {\r\n\t\tdocument.getElementsByClassName('loading')[0].style.display = 'inline-block';\r\n\t\tdocument.getElementsByClassName('loading')[1].style.display = 'inline-block';\r\n\t\taxios\r\n\t\t\t.post('/uploadWeb/' + dnsNameInputValue)\r\n\t\t\t.then((operation) => {\r\n\t\t\t\tdocument.getElementsByClassName('loading')[0].style.display = 'none';\r\n\t\t\t\tdocument.getElementsByClassName('loading')[1].style.display = 'none';\r\n\t\t\t\tsuccessMessage('Contract deployed to address ' + operation.data.address);\r\n\t\t\t})\r\n\t\t\t.catch((e) => {\r\n\t\t\t\terrorAlert(e.response.data.code);\r\n\t\t\t});\r\n\t}\r\n}\r\n\r\nfunction tableInsert(resp, count) {\r\n\tconst tBody = document.getElementById('website-deployers-table').getElementsByTagName('tbody')[0];\r\n\tconst row = tBody.insertRow(-1);\r\n\tconst url = 'http://' + resp.name + '.massa/';\r\n\r\n\tconst cell0 = row.insertCell();\r\n\tconst cell1 = row.insertCell();\r\n\tconst cell2 = row.insertCell();\r\n\tconst cell3 = row.insertCell();\r\n\r\n\tcell0.innerHTML = resp.name;\r\n\tcell1.innerHTML = resp.address;\r\n\tcell2.innerHTML = \"<a href='\" + url + \"'>\" + url + '</a>';\r\n\tcell3.innerHTML =\r\n\t\t\"<div><input id='fileid\" +\r\n\t\tcount +\r\n\t\t\"' type='file' hidden/><button id='upload-website\" +\r\n\t\tcount +\r\n\t\t\"'\" +\r\n\t\t\"class='primary-button' id='buttonid' type='button' value='Upload MB' >Upload</button><span style='display: none' class='spinner-border loading\" +\r\n\t\tcount +\r\n\t\t\"' role='status'><img src='./logo.png' class='massa-logo-spinner' alt='Massa logo' /></span></div> \";\r\n\r\n\tdocument.getElementById(`upload-website${count}`).addEventListener('click', function () {\r\n\t\tdocument.getElementById(`fileid${count}`).value = null;\r\n\t\tdocument.getElementById(`fileid${count}`).click();\r\n\t});\r\n\r\n\tdocument.getElementById(`fileid${count}`).addEventListener('change', function (evt) {\r\n\t\tlet files = evt.target.files; // get files\r\n\t\tlet f = files[0];\r\n\t\tuploadWebsite(f, count);\r\n\t});\r\n}\r\n\r\nfunction uploadWebsite(file, count) {\r\n\tconst bodyFormData = new FormData();\r\n\tbodyFormData.append('zipfile', file);\r\n\tdocument.getElementsByClassName('loading' + count)[0].style.display = 'inline-block';\r\n\taxios({\r\n\t\turl: `/fillWeb/${deployers[count].address}`,\r\n\t\tmethod: 'POST',\r\n\t\tdata: bodyFormData,\r\n\t\theaders: {\r\n\t\t\t'Content-Type': 'multipart/form-data',\r\n\t\t},\r\n\t})\r\n\t\t.then((operation) => {\r\n\t\t\tdocument.getElementsByClassName('loading' + count)[0].style.display = 'none';\r\n\t\t\tsuccessMessage('Website uploaded with operation ID : ' + operation);\r\n\t\t})\r\n\t\t.catch((e) => {\r\n\t\t\terrorAlert(e.response.data.code);\r\n\t\t});\r\n}\r\n"
