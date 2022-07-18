DELETE FROM USER_INFO WHERE USER_ID="001";
DELETE FROM ATTENDANCE_INFO WHERE USER_ID="001";

INSERT INTO USER_INFO(
    USER_ID
    ,EMAIL_ADDRESS
    ,PASSWORD
    ,CREATE_USER
    ,UPDATE_USER
    ,CREATE_DATE_TIME
    ,UPDATE_DATE_TIME
    )
    VALUES(
    "001",
    "test@test.com",
    "password",
    "test",
    "test",
    now(),
    now()
    );

INSERT INTO ATTENDANCE_INFO(
    USER_ID
    ,SCHEDULED_ATTENDANCE_DATE
    ,SCHEDULED_ATTENDANCE_TIME
    ,SCHEDULED_LEAVING_TIME
    ,ACHIEVED_ATTENDANCE_TIME
    ,ACHIEVED_LEAVING_TIME
    ,ATTEND_FLG
    ,LEAVE_FLG
    ,CREATE_USER
    ,UPDATE_USER
    ,CREATE_DATE_TIME
    ,UPDATE_DATE_TIME
    )
    VALUES(
    "001",
    "2022/7/16",
    "10:00",
    "19:00",
    null,
    null,
    0,
    0,
    "test",
    "test",
    now(),
    now()
    );
