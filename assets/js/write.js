$(function () {
	//Когда меняется инпкт поля контент мы шлем ajax запрос чтобы он нам отобразил html в #md_html
	$("#content").bind("input change", function () {
		$.post(
			"/gethtml",
			 {md: $("#content").val()},
			 function (response) {
			    $("#md_html").html(response.html)
		});
	});
})