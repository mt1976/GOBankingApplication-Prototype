USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCalenderHoliday]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCalenderHoliday]
AS
SELECT        {{!SQL.SOURCE}}.CountryHolidaysDates.InternalId AS ID, 'Holiday' AS Type, {{!SQL.SOURCE}}.CountryHolidaysDates.HDate AS Date, '' AS Time, {{!SQL.SOURCE}}.CountryHolidaysDates.Name, {{!SQL.SOURCE}}.Country.Name AS Description
FROM            {{!SQL.SOURCE}}.CountryHolidaysDates INNER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.CountryHolidaysDates.Code = {{!SQL.SOURCE}}.Country.Code
WHERE        ({{!SQL.SOURCE}}.CountryHolidaysDates.InternalDeleted IS NULL)
GO
