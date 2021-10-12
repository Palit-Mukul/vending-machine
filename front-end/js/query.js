<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.0/jquery.min.js"></script>
<script>
    $(document).ready(function () {
    $("#show-item").click(function (e) {
       $.ajax({
            type: "GET",
            url: "http://localhost:8009/itemlist",
            dataType: "json",
            success: function (result, status, xhr) {
                var table = $("<table><tr><th>Available Items</th></tr>");
                table.append("<tr><td>Name:</td><td>" + result["Name"] + "</td></tr>");
                table.append("<tr><td>Price:</td><td>" + result["Price"] + "</td></tr>");
                table.append("<tr><td>Count:</td><td>" + result["Count"] + "°C</td></tr>");
                $("#message").html(table);
            },
            error: function (xhr, status, error) {
                alert("Result: " + status + " " + error + " " + xhr.status + " " + xhr.statusText)
            }
        });
    });

    function Validate() {
    var errorMessage = "";
    if ($("#citySelect").val() == "Select") {
    errorMessage += "► Select City";
}
    return errorMessage;
}
});
</script>