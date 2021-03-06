require('material-design-lite/dist/material.indigo-blue.min.css');
require('material-design-lite/dist/material.min.js');
require('../css/app.css');
import hljs from 'highlight.js/lib/index.js';
import $ from 'jquery';
import 'jquery-ui/ui/core.js';
import 'jquery-ui/ui/widgets/datepicker.js';
import 'jquery-ui/themes/base/core.css';
import 'jquery-ui/themes/base/theme.css';
import 'jquery-ui/themes/base/datepicker.css';
require('highlight.js/styles/solarized-dark.css');
import Chart from 'chart.js/dist/Chart.min.js';

$(function() {
    hljs.initHighlightingOnLoad();

    $(".datepicker").datepicker({
        dateFormat: "yy/mm/dd"
    });

    $("button.clear").on('click', () => {
        $("input, textarea").val("");
    });

    $("button.regist-blog").on('click', () => {
        sendWriteData();
    });

    $("button.good").on('click', () => {
        incrementGood();
    });

    getRightColumn();
});

const sendWriteData = () => {
    if (confirm("本当に登録してもよろしいですか？")) {
        var params = $('form[name=form-blog]').serialize();

        $.post("/regist_blog/", params, (result) => {
            if (result === "success") {
                alert("登録しました");

                $("button.clear").trriger("click");
            }
        });
    }
};

const incrementGood = () => {
    var id = $("#blog_id").val();

    $.get("/increment_good/", {
        blog_id: id
    }, (result) => {
        $(".good").prop("disabled", true);
        $(".fab-count").text(result);
    });
};

const getRightColumn = () => {
    if ($("#right-column").length < 1) {
        return;
    }

    $.get("/popular_list/", (result) => {
        $("#right-column").empty();

        console.log(result.PopularList);
        $.each(result.PopularList, (i, val) => {
            $("#right-column").append($('<li class="mdl-list__item"></li>').append($('<a href="/detail/' + val.ID + '/" class="mdl-list__item-secondary-action"></a>').text(val.Title + " いいね:" + val.BlogCount.GoodCount)));
        });
    });

    categoryChart();
}

const categoryChart = () => {
    if ($("#right-column").length < 1) {
        return;
    }

    $.get("/category_chart/", (result) => {
        var ctx = document.getElementById("category-chart");
        var labels = [];
        var datas = [];

        $.each(result.CategoryCounts, (idx, value) => {
            labels.push(value.Name);
            datas.push(value.CategoryCount);
        });

        new Chart(ctx, {
            type: 'pie',
            data: {
                datasets: [{
                    data: datas,
                    backgroundColor: [
                        "#FF6384",
                        "#36A2EB",
                        "#FFCE56",
                        "#98FB98",
                    ],
                    label: "投稿カテゴリ"
                }],
                labels: labels
            },
        });
    });
}
