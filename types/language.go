package types

type LanguageType struct {
  Short            string
  Pluralizations   RulesType
  Singularizations RulesType
  Irregulars       IrregularsType
  Uncountables     UncountablesType
}

func (self *LanguageType) Pluralize(str string) string {
  if self.Uncountables.Contains(str) {
    return str
  } else if opposite, ok := self.Irregulars[str]; ok {
    return opposite
  } else {
    for _, rule := range self.Pluralizations {
      if rule.Regexp.MatchString(str) {
        return rule.Regexp.ReplaceAllString(str, rule.Replacer)
      }
    }
  }

  return str
}

func (self *LanguageType) Singularize(str string) string {
  if self.Uncountables.Contains(str) {
    return str
  } else if opposite, ok := self.Irregulars[str]; ok {
    return opposite
  } else {
    for _, rule := range self.Singularizations {
      if rule.Regexp.MatchString(str) {
        return rule.Regexp.ReplaceAllString(str, rule.Replacer)
      }
    }
  }

  return str
}

func (self *LanguageType) Plural(matcher, replacer string) *LanguageType {
  self.Pluralizations = append(self.Pluralizations, Rule(matcher, replacer))

  return self
}

func (self *LanguageType) Singular(matcher, replacer string) *LanguageType {
  self.Singularizations = append(self.Singularizations, Rule(matcher, replacer))

  return self
}

func (self *LanguageType) Irregular(singlular, plural string) *LanguageType {
  self.Irregulars[singlular] = plural
  self.Irregulars[plural] = singlular

  return self
}

func (self *LanguageType) Uncountable(uncountable string) *LanguageType {
  self.Uncountables = append(self.Uncountables, uncountable)

  return self
}

func Language(short string) (language *LanguageType) {
  language = new(LanguageType)

  language.Pluralizations = make(RulesType, 0)
  language.Singularizations = make(RulesType, 0)
  language.Irregulars = make(IrregularsType)
  language.Uncountables = make(UncountablesType, 0)

  return
}
