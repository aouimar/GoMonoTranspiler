package translater_test

import (
	"testing"

	"aouimar.com/translater"
)

type TestData struct {
	input    string
	result   string
	hasError bool
}

func TestUseDatabse(t *testing.T) {

	datas := []TestData{
		{"USE [facebook]", `\c "facebook";`, false},
		{"USES [face~~~]", ``, true},
	}

	for _, data := range datas {
		result, err := translater.Translate(data.input)

		if data.hasError {
			if !err {
				t.Logf("PASSED : the parse doesn't succeed as expected with %s as a result", result)
			} else {
				t.Errorf("FAILED: this should not be parsed!! , we expected %s and got %s", data.result, result)
			}
		} else {
			if err && result != data.result {
				t.Errorf("FAILED: we expected %s and got %s", data.result, result)
			} else if !err && result == data.result {
				t.Logf("PASSED: we expected %s and got %s", data.result, result)
			} else {
				t.Errorf("FAILED: unExpected Scenario we expected %s and got %s", data.result, result)
			}
		}
	}
}

func TestTranslateWithMassiveValidInput(t *testing.T) {
	//Setup the input
	input :=
		`USE [user]
	 GO CREATE SCHEMA [user] GO CREATE TABLE [user].[BlackList Users](
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
	ALTER TABLE [user].[BlackList Users]  WITH CHECK ADD  CONSTRAINT [Fk_BlackList Users_Persons] FOREIGN KEY([IdBlackList])
	REFERENCES [user].[Person] ([IdPerson])
	GO
	ALTER TABLE [user].[BlackList Users] CHECK CONSTRAINT [Fk_BlackList Users_Persons]
	GO`

	str, err := translater.Translate(input)
	if err {
		t.Errorf("%s", str)
	} else {
		t.Logf("%s", str)
	}

}
