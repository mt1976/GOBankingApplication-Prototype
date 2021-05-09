/****** Object:  View [{{!SQL.SCHEMA}}].[sienaMandatedUser]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaMandatedUser]
AS
SELECT        {{!SQL.SOURCE}}.MandatedUser.MandatedUserKeyCounterpartyFirm, {{!SQL.SOURCE}}.MandatedUser.MandatedUserKeyCounterpartyCentre, {{!SQL.SOURCE}}.MandatedUser.MandatedUserKeyUserName, {{!SQL.SOURCE}}.MandatedUser.TelephoneNumber, 
                         {{!SQL.SOURCE}}.MandatedUser.EmailAddress, {{!SQL.SOURCE}}.MandatedUser.Active, {{!SQL.SOURCE}}.MandatedUser.FirstName, {{!SQL.SOURCE}}.MandatedUser.Surname, ISNULL({{!SQL.SOURCE}}.MandatedUser.DateOfBirth, CAST('1900-01-01' AS DATE)) AS DateOfBirth, 
                         {{!SQL.SOURCE}}.MandatedUser.Postcode, {{!SQL.SOURCE}}.MandatedUser.NationalIDNo, {{!SQL.SOURCE}}.MandatedUser.PassportNo, {{!SQL.SOURCE}}.MandatedUser.Country, {{!SQL.SOURCE}}.Country.Name AS CountryName, {{!SQL.SOURCE}}.Firm.FullName AS FirmName, 
                         {{!SQL.SOURCE}}.Centre.FullName AS CentreName, {{!SQL.SOURCE}}.MandatedUser.Notify, {{!SQL.SOURCE}}.MandatedUser.SystemUser
FROM            {{!SQL.SOURCE}}.MandatedUser INNER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.MandatedUser.Country = {{!SQL.SOURCE}}.Country.Code INNER JOIN
                         {{!SQL.SOURCE}}.Firm ON {{!SQL.SOURCE}}.MandatedUser.MandatedUserKeyCounterpartyFirm = {{!SQL.SOURCE}}.Firm.FirmName INNER JOIN
                         {{!SQL.SOURCE}}.Centre ON {{!SQL.SOURCE}}.MandatedUser.MandatedUserKeyCounterpartyCentre = {{!SQL.SOURCE}}.Centre.ShortName
WHERE        ({{!SQL.SOURCE}}.MandatedUser.InternalDeleted IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[28] 4[24] 2[24] 3) )"
      End
      Begin PaneConfiguration = 1
         NumPanes = 3
         Configuration = "(H (1 [50] 4 [25] 3))"
      End
      Begin PaneConfiguration = 2
         NumPanes = 3
         Configuration = "(H (1 [50] 2 [25] 3))"
      End
      Begin PaneConfiguration = 3
         NumPanes = 3
         Configuration = "(H (4 [30] 2 [40] 3))"
      End
      Begin PaneConfiguration = 4
         NumPanes = 2
         Configuration = "(H (1 [56] 3))"
      End
      Begin PaneConfiguration = 5
         NumPanes = 2
         Configuration = "(H (2 [66] 3))"
      End
      Begin PaneConfiguration = 6
         NumPanes = 2
         Configuration = "(H (4 [50] 3))"
      End
      Begin PaneConfiguration = 7
         NumPanes = 1
         Configuration = "(V (3))"
      End
      Begin PaneConfiguration = 8
         NumPanes = 3
         Configuration = "(H (1[56] 4[18] 2) )"
      End
      Begin PaneConfiguration = 9
         NumPanes = 2
         Configuration = "(H (1 [75] 4))"
      End
      Begin PaneConfiguration = 10
         NumPanes = 2
         Configuration = "(H (1[66] 2) )"
      End
      Begin PaneConfiguration = 11
         NumPanes = 2
         Configuration = "(H (4 [60] 2))"
      End
      Begin PaneConfiguration = 12
         NumPanes = 1
         Configuration = "(H (1) )"
      End
      Begin PaneConfiguration = 13
         NumPanes = 1
         Configuration = "(V (4))"
      End
      Begin PaneConfiguration = 14
         NumPanes = 1
         Configuration = "(V (2))"
      End
      ActivePaneConfig = 0
   End
   Begin DiagramPane = 
      Begin Origin = 
         Top = -96
         Left = 0
      End
      Begin Tables = 
         Begin Table = "MandatedUser"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 207
               Right = 328
            End
            DisplayFlags = 280
            TopColumn = 18
         End
         Begin Table = "Country"
            Begin Extent = 
               Top = 6
               Left = 366
               Bottom = 136
               Right = 571
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Firm"
            Begin Extent = 
               Top = 140
               Left = 415
               Bottom = 270
               Right = 620
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Centre"
            Begin Extent = 
               Top = 210
               Left = 38
               Bottom = 340
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
      End
   End
   Begin SQLPane = 
   End
   Begin DataPane = 
      Begin ParameterDefaults = ""
      End
      Begin ColumnWidths = 20
         Width = 284
         Width = 2520
         Width = 1500
         Width = 3015
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 2445
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
      End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 1440
         Alias = 900
         Tabl' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaMandatedUser'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'e = 1170
         Output = 720
         Append = 1400
         NewValue = 1170
         SortType = 1350
         SortOrder = 1410
         GroupBy = 1350
         Filter = 1350
         Or = 1350
         Or = 1350
         Or = 1350
      End
   End
End
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaMandatedUser'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaMandatedUser'
GO
