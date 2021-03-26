USE [user]
GO
/****** Object:  User [user]    Script Date: 3/16/2021 11:33:21 PM ******/
CREATE USER [user] FOR LOGIN [user] WITH DEFAULT_SCHEMA=[user]
GO
/****** Object:  Schema [user]    Script Date: 3/16/2021 11:33:21 PM ******/
CREATE SCHEMA [user]
GO
/****** Object:  Table [user].[BlackList Users]    Script Date: 3/16/2021 11:33:21 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [user].[BlackList Users](
	[IdBlackList] [varchar](20) NOT NULL,
	[Description] [varchar](200) NULL,
 CONSTRAINT [Pk_BlackList Users_IdBlackList] PRIMARY KEY CLUSTERED 
(
	[IdBlackList] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [user].[Driver]    Script Date: 3/16/2021 11:33:21 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [user].[Driver](
	[IdDriver] [varchar](20) NOT NULL,
	[FirstName] [varchar](20) NULL,
	[SecondName] [varchar](20) NULL,
	[Email] [varchar](30) NULL,
	[Geolocation] [geography] NULL,
 CONSTRAINT [Pk_Driver_IdDriver] PRIMARY KEY CLUSTERED 
(
	[IdDriver] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [user].[GPoint]    Script Date: 3/16/2021 11:33:21 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [user].[GPoint](
	[IdPoint] [int] NOT NULL,
	[PathId] [int] NULL,
	[GeoPoint] [geography] NULL,
 CONSTRAINT [Pk_Point_IdPoint] PRIMARY KEY CLUSTERED 
(
	[IdPoint] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [user].[Path]    Script Date: 3/16/2021 11:33:21 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [user].[Path](
	[Idpath] [int] NOT NULL,
	[PersonId] [varchar](20) NULL,
	[NamePath] [varchar](100) NULL,
	[ComputedTarif] [money] NULL,
	[DriverId] [varchar](20) NULL,
	[PositionDeparture] [geography] NULL,
	[PositionArrival] [geography] NULL,
 CONSTRAINT [Pk_Path_Idpath] PRIMARY KEY CLUSTERED 
(
	[Idpath] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [user].[Payment Method]    Script Date: 3/16/2021 11:33:21 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [user].[Payment Method](
	[IdPaymentMethod] [varchar](20) NOT NULL,
	[Kind] [varchar](20) NULL,
	[IdPP] [varchar](20) NULL,
 CONSTRAINT [Pk_Payment Method_IdPaymentMethod] PRIMARY KEY CLUSTERED 
(
	[IdPaymentMethod] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [user].[Person]    Script Date: 3/16/2021 11:33:21 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [user].[Person](
	[IdPerson] [varchar](20) NOT NULL,
	[FirstName] [varchar](20) NULL,
	[SecondName] [varchar](20) NULL,
	[Email] [varchar](30) NULL,
	[Adress] [varchar](50) NULL,
	[Country] [char](2) NULL,
	[City] [varchar](20) NULL,
	[Geolocation] [geography] NULL,
 CONSTRAINT [Pk_Persons_IdPerson] PRIMARY KEY CLUSTERED 
(
	[IdPerson] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [user].[PreEstimatedPoint]    Script Date: 3/16/2021 11:33:21 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [user].[PreEstimatedPoint](
	[IdPreEstimatedP] [int] NOT NULL,
	[PathId] [int] NULL,
	[GeoPoint] [geography] NULL,
 CONSTRAINT [Pk_PreEstimatedPoint_IdPreEstimatedP] PRIMARY KEY CLUSTERED 
(
	[IdPreEstimatedP] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
ALTER TABLE [user].[BlackList Users]  WITH CHECK ADD  CONSTRAINT [Fk_BlackList Users_Persons] FOREIGN KEY([IdBlackList])
REFERENCES [user].[Person] ([IdPerson])
GO
ALTER TABLE [user].[BlackList Users] CHECK CONSTRAINT [Fk_BlackList Users_Persons]
GO
ALTER TABLE [user].[GPoint]  WITH CHECK ADD  CONSTRAINT [Fk_Point_Path] FOREIGN KEY([PathId])
REFERENCES [user].[Path] ([Idpath])
GO
ALTER TABLE [user].[GPoint] CHECK CONSTRAINT [Fk_Point_Path]
GO
ALTER TABLE [user].[Path]  WITH CHECK ADD  CONSTRAINT [Fk_Path_Driver] FOREIGN KEY([DriverId])
REFERENCES [user].[Driver] ([IdDriver])
GO
ALTER TABLE [user].[Path] CHECK CONSTRAINT [Fk_Path_Driver]
GO
ALTER TABLE [user].[Path]  WITH CHECK ADD  CONSTRAINT [Fk_Path_Persons] FOREIGN KEY([PersonId])
REFERENCES [user].[Person] ([IdPerson])
GO
ALTER TABLE [user].[Path] CHECK CONSTRAINT [Fk_Path_Persons]
GO
ALTER TABLE [user].[Payment Method]  WITH CHECK ADD  CONSTRAINT [Fk_Payment Method_Person] FOREIGN KEY([IdPP])
REFERENCES [user].[Person] ([IdPerson])
GO
ALTER TABLE [user].[Payment Method] CHECK CONSTRAINT [Fk_Payment Method_Person]
GO
ALTER TABLE [user].[PreEstimatedPoint]  WITH CHECK ADD  CONSTRAINT [Fk_PreEstimatedPoint_Path] FOREIGN KEY([PathId])
REFERENCES [user].[Path] ([Idpath])
GO
ALTER TABLE [user].[PreEstimatedPoint] CHECK CONSTRAINT [Fk_PreEstimatedPoint_Path]
GO
