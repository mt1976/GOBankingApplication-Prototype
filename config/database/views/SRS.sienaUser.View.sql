USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaUser]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaUser] AS SELECT        {{!SQL.SOURCE}}.Usr.UserName, {{!SQL.SOURCE}}.Usr.FullName, {{!SQL.SOURCE}}.Usr.Type, {{!SQL.SOURCE}}.Usr.TradingEntity, {{!SQL.SOURCE}}.UserDetailsSettings.DefaultEnterBook, {{!SQL.SOURCE}}.UserDetailsSettings.EmailAddress, {{!SQL.SOURCE}}.UserDetailsSettings.Enabled,
                         {{!SQL.SOURCE}}.UserDetailsSettings.ExternalUserIds, {{!SQL.SOURCE}}.UserDetailsSettings.Language, {{!SQL.SOURCE}}.UserDetailsSettings.LocalCurrency, {{!SQL.SOURCE}}.UserDetailsSettings.Role, {{!SQL.SOURCE}}.UserDetailsSettings.TelephoneNumber,
                         {{!SQL.SOURCE}}.UserDetailsSettings.TokenId, {{!SQL.SOURCE}}.UserDetailsSettings.TradingEntity AS Entity, {{!SQL.SOURCE}}.UserDetailsSettings.UserCode
FROM            {{!SQL.SOURCE}}.Usr INNER JOIN
                         {{!SQL.SOURCE}}.UserDetailsSettings ON {{!SQL.SOURCE}}.Usr.UserName = {{!SQL.SOURCE}}.UserDetailsSettings.UserName
WHERE        ({{!SQL.SOURCE}}.Usr.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.UserDetailsSettings.InternalDeleted IS NULL)


GO
