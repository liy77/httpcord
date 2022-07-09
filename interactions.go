package httpcord

import (
	"github.com/JustAWaifuHunter/httpcord/permissions"
)

type (
	InteractionType              int
	InteractionCallbackType      int
	ApplicationCommandType       int
	ButtonStyle                  int
	TextStyle                    int
	ComponentType                int
	ApplicationCommandOptionType int
	AnyComponent                 interface{}
)

// Interaction Types

const (
	PingInteraction InteractionType = iota + 1
	ApplicationCommandInteraction
	MessageComponentInteraction
	AutoCompleteInteraction
	ModalSubmitInteraction
)

// Interaction Callback Types

const (
	PongResponse InteractionCallbackType = iota + 1
	_
	_
	ChannelMessageWithSourceResponse
	DeferredChannelMessageWithSourceResponse
	DeferredUpdateResponse
	UpdateMessageResponse
	ApplicationCommandAutoCompleteResultResponse
	ModalResponse
)

// Commands Types

const (
	ChatInputApplicationCommandType ApplicationCommandType = iota + 1
	UserApplicationCommandType
	MessageApplicationCommandType
)

// Button Styles

const (
	PrimaryButtonStyle ButtonStyle = iota + 1
	SecondaryButtonStyle
	SuccessButtonStyle
	DangerButtonStyle
	LinkButtonStyle
)

// Text Styles

const (
	ShortTextStyle TextStyle = iota + 1
	ParagraphTextStyle
)

// Component Types

const (
	ActionRowComponentType ComponentType = iota + 1
	ButtonComponentType
	SelectMenuComponentType
	InputTextComponentType
)

// Application Command Option Type

const (
	SubCommandApplicationCommandOptionType ApplicationCommandOptionType = iota + 1
	SubCommandGroupApplicationCommandOptionType
	StringApplicationCommandOptionType
	IntApplicationCommandOptionType
	BoolApplicationCommandOptionType
	UserApplicationCommandOptionType
	ChannelApplicationCommandOptionType
	RoleApplicationCommandOptionType
	MentionableApplicationCommandOptionType
	NumberApplicationCommandOptionType
	AttachmentApplicationCommandOptionType
)

type Interaction struct {
	ID            Snowflake       `json:"id"`
	ApplicationID Snowflake       `json:"application_id"`
	Type          InteractionType `json:"type"`
	Data          interface{}     `json:"data,omitempty"`
	GuildID       Snowflake       `json:"guild_id"`
	ChannelID     Snowflake       `json:"channel_id"`
	Member        *Member         `json:"member"`
	User          *User           `json:"user"`
	Token         string          `json:"token"`
	Message       *Message        `json:"message,omitempty"`
	Version       int             `json:"version,omitempty"`
	Locale        string          `json:"locale"`
	GuildLocale   string          `json:"guild_locale"`
}

type ApplicationCommandInteractionData struct {
	ID       Snowflake                    `json:"id"`
	Name     string                       `json:"name"`
	Type     ApplicationCommandOptionType `json:"type"`
	Resolved ResolvedData                 `json:"resolved,omitempty"`
	Options  []ApplicationCommandOption   `json:"options,omitempty"`
	GuildID  Snowflake                    `json:"guild_id,omitempty"`
	TargetID Snowflake                    `json:"target_id,omitempty"`
}

type ResolvedData struct {
	Users       []*User       `json:"users,omitempty"`
	Members     []*Member     `json:"members,omitempty"`
	Roles       []*Role       `json:"roles,omitempty"`
	Channels    []*Channel    `json:"channels,omitempty"`
	Messages    []*Message    `json:"messages,omitempty"`
	Attachments []*Attachment `json:"attachment,omitempty"`
}

type InteractionCallbackData struct {
	TTS             bool                              `json:"tts,omitempty"`
	Content         string                            `json:"content,omitempty"`
	Embeds          []*Embed                          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions                  `json:"allowed_mentions,omitempty"`
	Flags           MessageFlag                       `json:"flags,omitempty"`
	Components      []*ActionRowComponent             `json:"components"`
	Attachments     []*Attachment                     `json:"attachments,omitempty"`
	Files           []*DiscordFile                    `json:"-"`
	Choices         []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
	CustomID        string                            `json:"custom_id,omitempty"`
	Title           string                            `json:"title,omitempty"`
}

type InteractionResponse struct {
	Type InteractionCallbackType  `json:"type"`
	Data *InteractionCallbackData `json:"data,omitempty"`
}

type TextInputComponent struct {
	Type        ComponentType `json:"type"`
	CustomID    string        `json:"custom_id,omitempty"`
	Style       TextStyle     `json:"style,omitempty"`
	Label       string        `json:"label,omitempty"`
	MinLength   *int          `json:"min_length,omitempty"`
	MaxLength   *int          `json:"max_length,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Value       string        `json:"value,omitempty"`
	Placeholder string        `json:"placeholder,omitempty"`
}

type ButtonComponent struct {
	Type     ComponentType `json:"type"`
	CustomID string        `json:"custom_id,omitempty"`
	Style    ButtonStyle   `json:"style,omitempty"`
	Label    string        `json:"label,omitempty"`
	Emoji    *Emoji        `json:"emoji,omitempty"`
	URL      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
}

type SelectMenuComponent struct {
	Type        ComponentType      `json:"type"`
	CustomID    string             `json:"custom_id"`
	Options     []*ComponentOption `json:"options,omitempty"`
	Placeholder string             `json:"placeholder,omitempty"`
	MinValues   *int               `json:"min_values,omitempty"`
	MaxValues   *int               `json:"max_values,omitempty"`
	Disabled    bool               `json:"disabled,omitempty"`
}

type ActionRowComponent struct {
	Type       ComponentType  `json:"type"`
	Components []AnyComponent `json:"components"`
}

type Modal struct {
	CustomID   string                `json:"custom_id"`
	Title      string                `json:"title"`
	Components []*ActionRowComponent `json:"components"`
}

type ComponentOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"`
	Default     bool   `json:"default"`
}

type ApplicationCommandOptionChoice struct {
	Name              string      `json:"name"`
	NameLocalizations Dictionary  `json:"name_localizations,omitempty"`
	Value             interface{} `json:"value"`
}

type ApplicationCommandOption struct {
	Type         ApplicationCommandOptionType     `json:"type"`
	Name         string                           `json:"name"`
	Description  string                           `json:"description"`
	Required     bool                             `json:"required,omitempty"`
	Choices      []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options      []ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes []ChannelType                    `json:"channel_types,omitempty"`
	MinValue     float64                          `json:"min_value,omitempty"`
	MaxValue     *float64                         `json:"max_value,omitempty"`
	Autocomplete bool                             `json:"autocomplete,omitempty"`
	Value        interface{}
}

type ApplicationCommand struct {
	// ID is the unique id of the command
	ID Snowflake `json:"id,omitempty"`
	// Type is	the type of command, defaults 1 if not set
	Type *ApplicationCommandType `json:"type,omitempty"`
	// Application ID is the unique id of the parent application
	ApplicationID Snowflake `json:"application_id,omitempty"`
	// GuildID guild id of the command, if not global
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// Name is a 1-32 character name
	Name string `json:"name"`
	// Localization dictionary for name field. Values follow the same restrictions as name
	NameLocalizations Dictionary `json:"name_localizations,omitempty"`
	// Description is a 1-100 character description for CHAT_INPUT commands, empty string for USER and MESSAGE commands
	Description string `json:"description,omitempty"`
	// Localization dictionary for description field. Values follow the same restrictions as description
	DescriptionLocalizations Dictionary `json:"description_localizations,omitempty"`
	// Options are the parameters for the command, max 25, only valid for CHAT_INPUT commands
	Options []ApplicationCommandOption `json:"options"`
	// Set of permissions represented as a bit set
	DefaultPermissions *permissions.PermissionBit `json:"default_member_permissions,omitempty"`
	// Indicates whether the command is available in DMs with the app, only for globally-scoped commands. By default, commands are visible.
	AllowUseInDMs *bool `json:"dm_permission,omitempty"`
	// DefaultPermission is whether the command is enabled by default when the app is added to a guild
	DefaultPermission *bool `json:"default_permission,omitempty"`
	// Version is an autoincrementing version identifier updated during substantial record changes
	Version Snowflake `json:"version,omitempty"`
}

type ComponentInteractionData struct {
	CustomID      string        `json:"custom_id"`
	ComponentType ComponentType `json:"component_type"`
	Values        []string      `json:"values"`
}

type ModalSubmitInteractionData struct {
	CustomID   string                `json:"custom_id"`
	Components []*ActionRowComponent `json:"components"`
}

func (i *Interaction) ModalSubmitData() ModalSubmitInteractionData {
	if i.Type != ModalSubmitInteraction {
		panic("The Interaction is not a ModalSubmit")
	}

	return i.Data.(ModalSubmitInteractionData)
}

func (i *Interaction) ApplicationCommandData() ApplicationCommandInteractionData {
	if i.Type != ApplicationCommandInteraction {
		panic("The Interaction is not a ApplicationCommand")
	}

	return i.Data.(ApplicationCommandInteractionData)
}

func (i *Interaction) ComponentData() ComponentInteractionData {
	if i.Type != MessageComponentInteraction {
		panic("The Interaction is not a Component")
	}

	return i.Data.(ComponentInteractionData)
}

func (c *ApplicationCommandInteractionData) UserValue(name string, required bool) *User {
	return c.GetOption(name, UserApplicationCommandOptionType, required).(*User)
}

func (c *ApplicationCommandInteractionData) StringValue(name string, required bool) string {
	return c.GetOption(name, UserApplicationCommandOptionType, required).(string)
}

func (c *ApplicationCommandInteractionData) BoolValue(name string, required bool) bool {
	return c.GetOption(name, UserApplicationCommandOptionType, required).(bool)
}

func (c *ApplicationCommandInteractionData) MemberValue(name string, required bool) *Member {
	return c.GetOption(name, UserApplicationCommandOptionType, required).(*Member)
}

func (c *ApplicationCommandInteractionData) IntValue(name string, required bool) int {
	return c.GetOption(name, UserApplicationCommandOptionType, required).(int)
}

func (c *ApplicationCommandInteractionData) NumberValue(name string, required bool) float64 {
	return c.GetOption(name, UserApplicationCommandOptionType, required).(float64)
}

func (c *ApplicationCommandInteractionData) ChannelValue(name string, required bool) *Channel {
	return c.GetOption(name, ChannelApplicationCommandOptionType, required).(*Channel)
}

func (c *ApplicationCommandInteractionData) AttachmentValue(name string, required bool) *Attachment {
	return c.GetOption(name, AttachmentApplicationCommandOptionType, required).(*Attachment)
}

// Returns Member, User or Role
func (c *ApplicationCommandInteractionData) MentionableValue(name string, required bool) interface{} {
	return c.GetOption(name, MentionableApplicationCommandOptionType, required)
}

func (c *ApplicationCommandInteractionData) RoleValue(name string, required bool) *Role {
	return c.GetOption(name, UserApplicationCommandOptionType, required).(*Role)
}

func (c *ApplicationCommandInteractionData) SubCommand() ApplicationCommandOption {
	option := c.Options[0]

	if option.Type != SubCommandApplicationCommandOptionType {
		panic("Could not find SubCommand")
	}

	return option
}

func (c *ApplicationCommandInteractionData) SubCommandGroup() ApplicationCommandOption {
	option := c.Options[0]

	if option.Type != SubCommandApplicationCommandOptionType {
		panic("Could not find SubCommandGroup")
	}

	return option
}

func (c *ApplicationCommandInteractionData) GetOption(name string, Type ApplicationCommandOptionType, required bool) interface{} {
	for _, option := range c.Options {
		if option.Name == name && option.Type == Type {
			return option.Value
		}
	}

	if required {
		panic("Could not find any option of type " + Type.String())
	}

	return nil
}

func (t ApplicationCommandOptionType) String() string {
	switch t {
	case IntApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(Integer)"
		}
	case NumberApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(Float64)"
		}
	case BoolApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(Boolean)"
		}
	case UserApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(User)"
		}
	case RoleApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(Role)"
		}
	case AttachmentApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(Attachment)"
		}
	case ChannelApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(Channel)"
		}
	case MentionableApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(Mentionable)"
		}
	case SubCommandApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(SubCommand)"
		}
	case SubCommandGroupApplicationCommandOptionType:
		{
			return "ApplicationCommandOptionType(SubCommandGroup)"
		}
	default:
		{
			return "ApplicationCommandOptionType(String)"
		}
	}
}

// Builders
// SelectMenuComponentBuilder

func NewSelectMenuComponentBuilder() *SelectMenuComponent {
	return &SelectMenuComponent{Type: SelectMenuComponentType}
}
func (s *SelectMenuComponent) AddOption(option *ComponentOption) *SelectMenuComponent {
	s.Options = append(s.Options, option)
	return s
}

func (s *SelectMenuComponent) SetOptions(options ...*ComponentOption) *SelectMenuComponent {
	s.Options = options
	return s
}

func (s *SelectMenuComponent) SetPlaceholder(placeholder string) *SelectMenuComponent {
	s.Placeholder = placeholder
	return s
}

func (s *SelectMenuComponent) SetMinValues(MinValues *int) *SelectMenuComponent {
	s.MinValues = MinValues
	return s
}

func (s *SelectMenuComponent) SetMaxValues(MaxValues *int) *SelectMenuComponent {
	s.MaxValues = MaxValues
	return s
}

func (s *SelectMenuComponent) IsDisabled(disabled bool) *SelectMenuComponent {
	s.Disabled = disabled
	return s
}

func (s *SelectMenuComponent) SetCustomID(customID string) *SelectMenuComponent {
	s.CustomID = customID
	return s
}

// ApplicationCommandOptionBuilder

func ApplicationCommandOptionBuilder() *ApplicationCommandOption {
	return &ApplicationCommandOption{}
}

func (o *ApplicationCommandOption) SetType(Type ApplicationCommandOptionType) *ApplicationCommandOption {
	o.Type = Type
	return o
}

func (o *ApplicationCommandOption) SetName(name string) *ApplicationCommandOption {
	o.Name = name
	return o
}

func (o *ApplicationCommandOption) SetDescription(description string) *ApplicationCommandOption {
	o.Description = description
	return o
}

func (o *ApplicationCommandOption) AddOption(ApplicationCommandOption ApplicationCommandOption) *ApplicationCommandOption {
	o.Options = append(o.Options, ApplicationCommandOption)
	return o
}

func (o *ApplicationCommandOption) SetOptions(options ...ApplicationCommandOption) *ApplicationCommandOption {
	o.Options = options
	return o
}

func (o *ApplicationCommandOption) AddChoice(choice ApplicationCommandOptionChoice) *ApplicationCommandOption {
	o.Choices = append(o.Choices, choice)
	return o
}

func (o *ApplicationCommandOption) SetChoices(choices ...ApplicationCommandOptionChoice) *ApplicationCommandOption {
	o.Choices = choices
	return o
}

func (o *ApplicationCommandOption) IsRequired(required bool) *ApplicationCommandOption {
	o.Required = required
	return o
}

func (o *ApplicationCommandOption) SetMinValue(MinValue float64) *ApplicationCommandOption {
	o.MinValue = MinValue
	return o
}

func (o *ApplicationCommandOption) SetMaxValue(MaxValue *float64) *ApplicationCommandOption {
	o.MaxValue = MaxValue
	return o
}

// TextInputComponentBuilder

func NewTextInputComponentBuilder() *TextInputComponent {
	return &TextInputComponent{Type: InputTextComponentType}
}

func (t *TextInputComponent) SetCustomID(customID string) *TextInputComponent {
	t.CustomID = customID
	return t
}

func (t *TextInputComponent) SetStyle(style TextStyle) *TextInputComponent {
	t.Style = style
	return t
}

func (t *TextInputComponent) SetLabel(label string) *TextInputComponent {
	t.Label = label
	return t
}

func (t *TextInputComponent) SetMinLength(MinLength *int) *TextInputComponent {
	t.MinLength = MinLength
	return t
}

func (t *TextInputComponent) SetMaxLength(MaxLength *int) *TextInputComponent {
	t.MaxLength = MaxLength
	return t
}

func (t *TextInputComponent) IsRequired(required bool) *TextInputComponent {
	t.Required = required
	return t
}

func (t *TextInputComponent) SetPlaceholder(placeholder string) *TextInputComponent {
	t.Placeholder = placeholder
	return t
}

// ButtonComponentBuilder

func NewButtonComponentBuilder() *ButtonComponent {
	return &ButtonComponent{Type: ButtonComponentType}
}

func (b *ButtonComponent) SetStyle(style ButtonStyle) *ButtonComponent {
	b.Style = style
	return b
}

func (b *ButtonComponent) SetLabel(label string) *ButtonComponent {
	b.Label = label
	return b
}

func (b *ButtonComponent) SetEmoji(emoji *Emoji) *ButtonComponent {
	b.Emoji = emoji
	return b
}

func (b *ButtonComponent) SetCustomID(customID string) *ButtonComponent {
	b.CustomID = customID
	return b
}

func (b *ButtonComponent) SetURL(URL string) *ButtonComponent {
	b.URL = URL
	return b
}

func (b *ButtonComponent) IsDisabled(disabled bool) *ButtonComponent {
	b.Disabled = disabled
	return b
}

// ActionRowComponentBuilder

func NewActionRowComponentBuilder() *ActionRowComponent {
	return &ActionRowComponent{Type: ActionRowComponentType}
}

func (a *ActionRowComponent) AddComponent(component AnyComponent) *ActionRowComponent {
	a.Components = append(a.Components, component)
	return a
}

func (a *ActionRowComponent) SetComponents(components ...AnyComponent) *ActionRowComponent {
	a.Components = components
	return a
}

// ApplicationCommandOptionChoiceBuilder

func NewApplicationCommandOptionChoiceBuilder() *ApplicationCommandOptionChoice {
	return &ApplicationCommandOptionChoice{}
}

func (c *ApplicationCommandOptionChoice) SetName(name string) *ApplicationCommandOptionChoice {
	c.Name = name
	return c
}

func (c *ApplicationCommandOptionChoice) SetValue(value interface{}) *ApplicationCommandOptionChoice {
	c.Value = value
	return c
}

func (c *ApplicationCommandOptionChoice) SetNameLocalizations(NameLocalizations Dictionary) *ApplicationCommandOptionChoice {
	c.NameLocalizations = NameLocalizations
	return c
}

// ApplicationCommandBuilder

func NewCommandBuilder() *ApplicationCommand {
	return &ApplicationCommand{}
}

func (c *ApplicationCommand) SetType(Type ApplicationCommandType) *ApplicationCommand {
	c.Type = &Type
	return c
}

func (c *ApplicationCommand) SetName(name string) *ApplicationCommand {
	c.Name = name
	return c
}

func (c *ApplicationCommand) SetDescription(description string) *ApplicationCommand {
	c.Description = description
	return c
}

func (c *ApplicationCommand) AddOption(option ApplicationCommandOption) *ApplicationCommand {
	c.Options = append(c.Options, option)
	return c
}

func (c *ApplicationCommand) SetOptions(options ...ApplicationCommandOption) *ApplicationCommand {
	c.Options = options
	return c
}

func (c *ApplicationCommand) AllowDM(allow bool) *ApplicationCommand {
	c.AllowUseInDMs = &allow
	return c
}

func (c *ApplicationCommand) SetDefaultPermissions(bits *permissions.PermissionBit) *ApplicationCommand {
	c.DefaultPermissions = bits
	return c
}

func (c *ApplicationCommand) SetNameLocalizations(NameLocalizations Dictionary) *ApplicationCommand {
	c.NameLocalizations = NameLocalizations
	return c
}

func (c *ApplicationCommand) SetDescriptionLocalizations(DescriptionLocalizations Dictionary) *ApplicationCommand {
	c.DescriptionLocalizations = DescriptionLocalizations
	return c
}
