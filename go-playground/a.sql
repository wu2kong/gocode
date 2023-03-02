# 2.1 cohort   复购用户数
SELECT
DATE_FORMAT(co1.create_time,'%Y%m') AS 首单月份,
COUNT(DISTINCT u_pay.user_id) as 新增用户数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =0, co2.user_id, NULL)) AS M0复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =1, co2.user_id, NULL)) AS M1复购课人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =2, co2.user_id, NULL)) AS M2复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =3, co2.user_id, NULL)) AS M3复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =4, co2.user_id, NULL)) AS M4复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =5, co2.user_id, NULL)) AS M5复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =6, co2.user_id, NULL)) AS M6复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =7, co2.user_id, NULL)) AS M7复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =8, co2.user_id, NULL)) AS M8复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =9, co2.user_id, NULL)) AS M9复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =10, co2.user_id, NULL)) AS M10复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =11, co2.user_id, NULL)) AS M11复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =12, co2.user_id, NULL)) AS M12复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =13, co2.user_id, NULL)) AS M13复购人数,
COUNT(DISTINCT IF(TIMESTAMPDIFF(MONTH, DATE_FORMAT(co1.create_time, '%Y-%m-01'), DATE_FORMAT(co2.create_time, '%Y-%m-01')) =14, co2.user_id, NULL)) AS M14复购人数
FROM
    # 大单用户且累计付费>xxx
    (
        SELECT
            co.user_id,
            MIN(co.card_order_id) as fir_id
        FROM
            dd_card_order co
            INNER JOIN v_user u ON co.user_id = u.user_id
            AND u.is_test = 0
            AND u.is_inner_user = 0
            /*排除测试用户*/
        WHERE
            co.is_test = 0
            AND co.state = 1
        GROUP BY
            1
    ) AS u_pay
    INNER JOIN dd_card_order co1 ON co1.card_order_id = u_pay.fir_id
    /*首次付费的订单信息*/
    and co1.order_money BETWEEN 0
    AND 50 — —
    and co1.order_money > 1000
    LEFT JOIN dd_card_order co2 ON co2.user_id = u_pay.user_id
    AND co2.state = 1
    and co2.is_test = 0
    and co2.create_time > co1.create_time
    and co2.order_money > 50
    /*大于50元的复购*/
GROUP BY
    1
HAVING
    首单月份 BETWEEN 201808
    AND 201910;