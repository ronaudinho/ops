package lepton

// CreateVolume is a stub to satisfy VolumeService interface
func (v *Vultr) CreateVolume(config *Config, name, label, data, size, provider string) error {
	return nil
}

// GetAllVolumes is a stub to satisfy VolumeService interface
func (v *Vultr) GetAllVolumes(config *Config) error {
	return nil
}

// UpdateVolume is a stub to satisfy VolumeService interface
func (v *Vultr) UpdateVolume(config *Config, name, label string) error {
	return nil
}

// DeleteVolume is a stub to satisfy VolumeService interface
func (v *Vultr) DeleteVolume(config *Config, name, label string) error {
	return nil
}

// AttachVolume is a stub to satisfy VolumeService interface
func (v *Vultr) AttachVolume(config *Config, image, name, label, mount string) error {
	return nil
}

// DetachVolume is a stub to satisfy VolumeService interface
func (v *Vultr) DetachVolume(config *Config, image, label string) error {
	return nil
}
