/****** Object:  View [{{!SQL.SCHEMA}}].[sienaAccountTransactions]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaAccountTransactions]
AS
SELECT        {{!SQL.SOURCE}}.DealLegsInfo.SienaReference, {{!SQL.SOURCE}}.DealLegsInfo.LegNo, {{!SQL.SOURCE}}.DealLegsInfo.MMLegNo, {{!SQL.SOURCE}}.DealLegsInfo.Narrative, {{!SQL.SOURCE}}.DealLegsInfo.Amount, {{!SQL.SOURCE}}.DealLegsInfo.StartInterestDate, {{!SQL.SOURCE}}.DealLegsInfo.EndInterestDate, 
                         {{!SQL.SOURCE}}.DealLegsInfo.Amortisation, {{!SQL.SOURCE}}.DealLegsInfo.InterestAmount, {{!SQL.SOURCE}}.DealLegsInfo.InterestAction, {{!SQL.SOURCE}}.DealLegsInfo.FixingDate, {{!SQL.SOURCE}}.DealLegsInfo.InterestCalculationDate, {{!SQL.SOURCE}}.DealLegsInfo.AmendmentAmount, 
                         {{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Currency.AmountDp
FROM            {{!SQL.SOURCE}}.Currency LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Deals ON {{!SQL.SOURCE}}.Currency.Code = {{!SQL.SOURCE}}.Deals.DealtCcy RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.DealLegsInfo ON {{!SQL.SOURCE}}.Deals.SienaReference = {{!SQL.SOURCE}}.DealLegsInfo.SienaReference
WHERE        ({{!SQL.SOURCE}}.Deals.DealType = 'Acct') AND ({{!SQL.SOURCE}}.Deals.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealLegsInfo.InternalDeleted IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[47] 4[12] 2[15] 3) )"
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
         Top = 0
         Left = 0
      End
      Begin Tables = 
         Begin Table = "DealLegsInfo"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 198
               Right = 263
            End
            DisplayFlags = 280
            TopColumn = 15
         End
         Begin Table = "Deals"
            Begin Extent = 
               Top = 6
               Left = 301
               Bottom = 214
               Right = 555
            End
            DisplayFlags = 280
            TopColumn = 14
         End
         Begin Table = "Currency"
            Begin Extent = 
               Top = 33
               Left = 574
               Bottom = 249
               Right = 946
            End
            DisplayFlags = 280
            TopColumn = 30
         End
      End
   End
   Begin SQLPane = 
   End
   Begin DataPane = 
      Begin ParameterDefaults = ""
      End
      Begin ColumnWidths = 38
         Width = 284
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 2700
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
    ' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaAccountTransactions'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'     Width = 1500
      End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 1440
         Alias = 900
         Table = 1170
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaAccountTransactions'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaAccountTransactions'
GO
